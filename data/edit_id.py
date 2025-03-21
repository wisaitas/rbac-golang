import json
import uuid

# อ่านไฟล์ JSON
with open('data/provinces.json', 'r', encoding='utf-8') as file:
    provinces = json.load(file)

with open('data/districts.json', 'r', encoding='utf-8') as file:
    districts = json.load(file)

with open('data/sub_districts.json', 'r', encoding='utf-8') as file:
    sub_districts = json.load(file)

with open('data/permission.json', 'r', encoding='utf-8') as file:
    permissions = json.load(file)

with open('data/roles.json', 'r', encoding='utf-8') as file:
    roles = json.load(file)

with open('data/users.json', 'r', encoding='utf-8') as file:
    users = json.load(file)

# แก้ไข id ให้เป็น UUI
for province in provinces:
    province['uuid'] = str(uuid.uuid4())

# สร้าง dictionary สำหรับเก็บความสัมพันธ์ระหว่าง province_id กับ province_uuid
province_id_to_uuid = {province['id']: province['uuid'] for province in provinces}

for district in districts:
    district['uuid'] = str(uuid.uuid4())
    # เพิ่ม field province_uuid โดยอ้างอิงจาก province_id
    if 'province_id' in district:
        district['province_uuid'] = province_id_to_uuid.get(district['province_id'])

# สร้าง dictionary สำหรับเก็บความสัมพันธ์ระหว่าง district_id กับ district_uuid
district_id_to_uuid = {district['id']: district['uuid'] for district in districts}

for sub_district in sub_districts:
    sub_district['uuid'] = str(uuid.uuid4())
    # เพิ่ม field district_uuid โดยอ้างอิงจาก district_id
    if 'district_id' in sub_district:
        sub_district['district_uuid'] = district_id_to_uuid.get(sub_district['district_id'])

for role in roles:
    role['id'] = str(uuid.uuid4())

for permission in permissions:
    permission['id'] = str(uuid.uuid4())

for user in users:
    user['id'] = str(uuid.uuid4())

# สร้างความสัมพันธ์ระหว่าง roles กับ permissions
roles_permissions = []
# สร้าง dictionary เพื่อหา permission_id จาก name
permission_name_to_id = {permission['name']: permission['id'] for permission in permissions}

# สำหรับ admin ให้มีทุก permissions
admin_role_id = next(role['id'] for role in roles if role['name'] == 'admin')
for permission in permissions:
    roles_permissions.append({
        "role_id": admin_role_id,
        "permission_id": permission['id']
    })

# บันทึกความสัมพันธ์ลงไฟล์
with open('data/roles_permissions.json', 'w', encoding='utf-8') as file:
    json.dump(roles_permissions, file, ensure_ascii=False, indent=2)

# สร้างความสัมพันธ์ระหว่าง users กับ roles
users_roles = []
# ค้นหา admin user
admin_user_id = next(user['id'] for user in users if user['username'] == 'admin')
# สร้างความสัมพันธ์ระหว่าง admin user กับ admin role
users_roles.append({
    "user_id": admin_user_id,
    "role_id": admin_role_id
})

# บันทึกความสัมพันธ์ลงไฟล์
with open('data/users-roles.json', 'w', encoding='utf-8') as file:
    json.dump(users_roles, file, ensure_ascii=False, indent=2)

# บันทึกกลับไปที่ไฟล์
with open('data/provinces.json', 'w', encoding='utf-8') as file:
    json.dump(provinces, file, ensure_ascii=False, indent=2)

with open('data/districts.json', 'w', encoding='utf-8') as file:
    json.dump(districts, file, ensure_ascii=False, indent=2)

with open('data/sub_districts.json', 'w', encoding='utf-8') as file:
    json.dump(sub_districts, file, ensure_ascii=False, indent=2)

with open('data/permission.json', 'w', encoding='utf-8') as file:
    json.dump(permissions, file, ensure_ascii=False, indent=2)

with open('data/roles.json', 'w', encoding='utf-8') as file:
    json.dump(roles, file, ensure_ascii=False, indent=2)

with open('data/users.json', 'w', encoding='utf-8') as file:
    json.dump(users, file, ensure_ascii=False, indent=2)

print('Edit ID Successfully')