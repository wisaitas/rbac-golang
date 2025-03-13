#!/bin/bash

# Directory paths
MODELS_DIR="internal/models"
REPOSITORIES_DIR="internal/repositories"
VALIDATES_DIR="internal/validates"

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

EOF
    else
        echo "Validate file already exists: $VALIDATE_FILE"
    fi
done

echo "Generation completed!"