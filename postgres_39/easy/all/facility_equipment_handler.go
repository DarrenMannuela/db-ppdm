package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_equipment

	for rows.Next() {
		var facility_equipment dto.Facility_equipment
		if err := rows.Scan(&facility_equipment.Facility_id, &facility_equipment.Facility_type, &facility_equipment.Equipment_id, &facility_equipment.Install_obs_no, &facility_equipment.Active_ind, &facility_equipment.Effective_date, &facility_equipment.Expiry_date, &facility_equipment.Install_date, &facility_equipment.Ppdm_guid, &facility_equipment.Remark, &facility_equipment.Remove_date, &facility_equipment.Remove_reason, &facility_equipment.Source, &facility_equipment.Use_description, &facility_equipment.Row_changed_by, &facility_equipment.Row_changed_date, &facility_equipment.Row_created_by, &facility_equipment.Row_created_date, &facility_equipment.Row_effective_date, &facility_equipment.Row_expiry_date, &facility_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_equipment to result
		result = append(result, facility_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityEquipment(c *fiber.Ctx) error {
	var facility_equipment dto.Facility_equipment

	setDefaults(&facility_equipment)

	if err := json.Unmarshal(c.Body(), &facility_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	facility_equipment.Row_created_date = formatDate(facility_equipment.Row_created_date)
	facility_equipment.Row_changed_date = formatDate(facility_equipment.Row_changed_date)
	facility_equipment.Effective_date = formatDateString(facility_equipment.Effective_date)
	facility_equipment.Expiry_date = formatDateString(facility_equipment.Expiry_date)
	facility_equipment.Install_date = formatDateString(facility_equipment.Install_date)
	facility_equipment.Remove_date = formatDateString(facility_equipment.Remove_date)
	facility_equipment.Row_effective_date = formatDateString(facility_equipment.Row_effective_date)
	facility_equipment.Row_expiry_date = formatDateString(facility_equipment.Row_expiry_date)






	rows, err := stmt.Exec(facility_equipment.Facility_id, facility_equipment.Facility_type, facility_equipment.Equipment_id, facility_equipment.Install_obs_no, facility_equipment.Active_ind, facility_equipment.Effective_date, facility_equipment.Expiry_date, facility_equipment.Install_date, facility_equipment.Ppdm_guid, facility_equipment.Remark, facility_equipment.Remove_date, facility_equipment.Remove_reason, facility_equipment.Source, facility_equipment.Use_description, facility_equipment.Row_changed_by, facility_equipment.Row_changed_date, facility_equipment.Row_created_by, facility_equipment.Row_created_date, facility_equipment.Row_effective_date, facility_equipment.Row_expiry_date, facility_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityEquipment(c *fiber.Ctx) error {
	var facility_equipment dto.Facility_equipment

	setDefaults(&facility_equipment)

	if err := json.Unmarshal(c.Body(), &facility_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_equipment.Facility_id = id

    if facility_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_equipment set facility_type = :1, equipment_id = :2, install_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, install_date = :7, ppdm_guid = :8, remark = :9, remove_date = :10, remove_reason = :11, source = :12, use_description = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where facility_id = :21")
	if err != nil {
		return err
	}

	facility_equipment.Row_changed_date = formatDate(facility_equipment.Row_changed_date)
	facility_equipment.Effective_date = formatDateString(facility_equipment.Effective_date)
	facility_equipment.Expiry_date = formatDateString(facility_equipment.Expiry_date)
	facility_equipment.Install_date = formatDateString(facility_equipment.Install_date)
	facility_equipment.Remove_date = formatDateString(facility_equipment.Remove_date)
	facility_equipment.Row_effective_date = formatDateString(facility_equipment.Row_effective_date)
	facility_equipment.Row_expiry_date = formatDateString(facility_equipment.Row_expiry_date)






	rows, err := stmt.Exec(facility_equipment.Facility_type, facility_equipment.Equipment_id, facility_equipment.Install_obs_no, facility_equipment.Active_ind, facility_equipment.Effective_date, facility_equipment.Expiry_date, facility_equipment.Install_date, facility_equipment.Ppdm_guid, facility_equipment.Remark, facility_equipment.Remove_date, facility_equipment.Remove_reason, facility_equipment.Source, facility_equipment.Use_description, facility_equipment.Row_changed_by, facility_equipment.Row_changed_date, facility_equipment.Row_created_by, facility_equipment.Row_effective_date, facility_equipment.Row_expiry_date, facility_equipment.Row_quality, facility_equipment.Facility_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchFacilityEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_equipment set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "install_date" || key == "remove_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where facility_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteFacilityEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_equipment dto.Facility_equipment
	facility_equipment.Facility_id = id

	stmt, err := db.Prepare("delete from facility_equipment where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_equipment.Facility_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


