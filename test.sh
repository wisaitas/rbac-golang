#!/bin/bash

# Directory paths
MODELS_DIR="internal/models"
REPOSITORIES_DIR="internal/repositories"
VALIDATES_DIR="internal/validates"
INITIAL_DIR="internal/initial"

# Create directories if they don't exist
mkdir -p "$REPOSITORIES_DIR"
mkdir -p "$VALIDATES_DIR"

# Function to convert snake_case to CamelCase
snake_to_camel() {
    echo "$1" | sed -r 's/(^|_)([a-z])/\U\2/g'
}

# Function to convert CamelCase to snake_case
camel_to_snake() {
    echo "$1" | sed -r 's/([A-Z])/_\L\1/g' | sed 's/^_//'
}

# Arrays to store model names for dependency injection
MODELS_ARRAY=()
MODELS_CAMEL_ARRAY=()

# Process each model file
for MODEL_FILE in "$MODELS_DIR"/*.go; do
    FILENAME=$(basename "$MODEL_FILE")
    
    # Skip base.go and context.go
    if [[ "$FILENAME" == "base.go" || "$FILENAME" == "context.go" ]]; then
        continue
    fi
    
    # Extract model name from filename (remove .go extension)
    MODEL_NAME=$(basename "$MODEL_FILE" .go)
    
    # Convert to CamelCase for struct name
    STRUCT_NAME=$(snake_to_camel "$MODEL_NAME")
    
    echo "Processing model: $STRUCT_NAME from $FILENAME"
    
    # Add to arrays for dependency injection
    MODELS_ARRAY+=("$MODEL_NAME")
    MODELS_CAMEL_ARRAY+=("$STRUCT_NAME")
    
    # Generate repository file
    REPO_FILE="$REPOSITORIES_DIR/$MODEL_NAME.go"
    
    if [ ! -f "$REPO_FILE" ]; then
        echo "Generating repository file: $REPO_FILE"
        
        cat > "$REPO_FILE" << EOF
package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type ${STRUCT_NAME}Repository interface {
	BaseRepository[models.${STRUCT_NAME}]
}

type ${MODEL_NAME}Repository struct {
	BaseRepository[models.${STRUCT_NAME}]
	db *gorm.DB
}

func New${STRUCT_NAME}Repository(db *gorm.DB, baseRepository BaseRepository[models.${STRUCT_NAME}]) ${STRUCT_NAME}Repository {
	return &${MODEL_NAME}Repository{
		BaseRepository: baseRepository,
		db:             db,
	}
}
EOF
    else
        echo "Repository file already exists: $REPO_FILE"
    fi
    
    # Generate validate file
    VALIDATE_FILE="$VALIDATES_DIR/$MODEL_NAME.go"
    
    if [ ! -f "$VALIDATE_FILE" ]; then
        echo "Generating validate file: $VALIDATE_FILE"
        
        cat > "$VALIDATE_FILE" << EOF
package validates

type ${STRUCT_NAME}Validate struct {
}

func New${STRUCT_NAME}Validate() *${STRUCT_NAME}Validate {
	return &${STRUCT_NAME}Validate{}
}
EOF
    else
        echo "Validate file already exists: $VALIDATE_FILE"
    fi
done

# Update dependency injection files
echo "Updating dependency injection files..."

# Update repositories.go
REPO_FILE="$INITIAL_DIR/repositorie.go"
REPO_STRUCT=""
REPO_INIT=""

for i in "${!MODELS_ARRAY[@]}"; do
    MODEL="${MODELS_ARRAY[$i]}"
    CAMEL="${MODELS_CAMEL_ARRAY[$i]}"
    
    REPO_STRUCT="${REPO_STRUCT}	${CAMEL}Repository    repositories.${CAMEL}Repository
"
    REPO_INIT="${REPO_INIT}		${CAMEL}Repository:    repositories.New${CAMEL}Repository(db, repositories.NewBaseRepository[models.${CAMEL}](db)),
"
done

# Create repositories.go file (without .new extension)
cat > "${REPO_FILE}" << EOF
package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
${REPO_STRUCT}}

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
${REPO_INIT}	}
}
EOF

# Update validates.go
VALIDATE_FILE="$INITIAL_DIR/validate.go"
VALIDATE_STRUCT=""
VALIDATE_INIT=""

for i in "${!MODELS_ARRAY[@]}"; do
    MODEL="${MODELS_ARRAY[$i]}"
    CAMEL="${MODELS_CAMEL_ARRAY[$i]}"
    
    VALIDATE_STRUCT="${VALIDATE_STRUCT}	${CAMEL}Validate    validates.${CAMEL}Validate
"
    VALIDATE_INIT="${VALIDATE_INIT}		${CAMEL}Validate:    *validates.New${CAMEL}Validate(),
"
done

# Create validate.go file (without .new extension)
cat > "${VALIDATE_FILE}" << EOF
package initial

import (
	"github.com/wisaitas/standard-golang/internal/validates"
)

type Validates struct {
${VALIDATE_STRUCT}}

func initializeValidates() *Validates {
	return &Validates{
${VALIDATE_INIT}	}
}
EOF

echo "Dependency injection files updated:"
echo "- ${REPO_FILE}"
echo "- ${VALIDATE_FILE}"
echo ""
echo "Generation completed!"