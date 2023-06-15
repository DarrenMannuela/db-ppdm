package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRMudPropertyRef(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_mud_property_ref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_mud_property_ref

	for rows.Next() {
		var r_mud_property_ref dto.R_mud_property_ref
		if err := rows.Scan(&r_mud_property_ref.Mud_property_ref, &r_mud_property_ref.Abbreviation, &r_mud_property_ref.Active_ind, &r_mud_property_ref.Effective_date, &r_mud_property_ref.Expiry_date, &r_mud_property_ref.Long_name, &r_mud_property_ref.Ppdm_guid, &r_mud_property_ref.Property_set_id, &r_mud_property_ref.Remark, &r_mud_property_ref.Short_name, &r_mud_property_ref.Source, &r_mud_property_ref.Row_changed_by, &r_mud_property_ref.Row_changed_date, &r_mud_property_ref.Row_created_by, &r_mud_property_ref.Row_created_date, &r_mud_property_ref.Row_effective_date, &r_mud_property_ref.Row_expiry_date, &r_mud_property_ref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_mud_property_ref to result
		result = append(result, r_mud_property_ref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRMudPropertyRef(c *fiber.Ctx) error {
	var r_mud_property_ref dto.R_mud_property_ref

	setDefaults(&r_mud_property_ref)

	if err := json.Unmarshal(c.Body(), &r_mud_property_ref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_mud_property_ref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_mud_property_ref.Row_created_date = formatDate(r_mud_property_ref.Row_created_date)
	r_mud_property_ref.Row_changed_date = formatDate(r_mud_property_ref.Row_changed_date)
	r_mud_property_ref.Effective_date = formatDateString(r_mud_property_ref.Effective_date)
	r_mud_property_ref.Expiry_date = formatDateString(r_mud_property_ref.Expiry_date)
	r_mud_property_ref.Row_effective_date = formatDateString(r_mud_property_ref.Row_effective_date)
	r_mud_property_ref.Row_expiry_date = formatDateString(r_mud_property_ref.Row_expiry_date)






	rows, err := stmt.Exec(r_mud_property_ref.Mud_property_ref, r_mud_property_ref.Abbreviation, r_mud_property_ref.Active_ind, r_mud_property_ref.Effective_date, r_mud_property_ref.Expiry_date, r_mud_property_ref.Long_name, r_mud_property_ref.Ppdm_guid, r_mud_property_ref.Property_set_id, r_mud_property_ref.Remark, r_mud_property_ref.Short_name, r_mud_property_ref.Source, r_mud_property_ref.Row_changed_by, r_mud_property_ref.Row_changed_date, r_mud_property_ref.Row_created_by, r_mud_property_ref.Row_created_date, r_mud_property_ref.Row_effective_date, r_mud_property_ref.Row_expiry_date, r_mud_property_ref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRMudPropertyRef(c *fiber.Ctx) error {
	var r_mud_property_ref dto.R_mud_property_ref

	setDefaults(&r_mud_property_ref)

	if err := json.Unmarshal(c.Body(), &r_mud_property_ref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_mud_property_ref.Mud_property_ref = id

    if r_mud_property_ref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_mud_property_ref set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, property_set_id = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where mud_property_ref = :18")
	if err != nil {
		return err
	}

	r_mud_property_ref.Row_changed_date = formatDate(r_mud_property_ref.Row_changed_date)
	r_mud_property_ref.Effective_date = formatDateString(r_mud_property_ref.Effective_date)
	r_mud_property_ref.Expiry_date = formatDateString(r_mud_property_ref.Expiry_date)
	r_mud_property_ref.Row_effective_date = formatDateString(r_mud_property_ref.Row_effective_date)
	r_mud_property_ref.Row_expiry_date = formatDateString(r_mud_property_ref.Row_expiry_date)






	rows, err := stmt.Exec(r_mud_property_ref.Abbreviation, r_mud_property_ref.Active_ind, r_mud_property_ref.Effective_date, r_mud_property_ref.Expiry_date, r_mud_property_ref.Long_name, r_mud_property_ref.Ppdm_guid, r_mud_property_ref.Property_set_id, r_mud_property_ref.Remark, r_mud_property_ref.Short_name, r_mud_property_ref.Source, r_mud_property_ref.Row_changed_by, r_mud_property_ref.Row_changed_date, r_mud_property_ref.Row_created_by, r_mud_property_ref.Row_effective_date, r_mud_property_ref.Row_expiry_date, r_mud_property_ref.Row_quality, r_mud_property_ref.Mud_property_ref)
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

func PatchRMudPropertyRef(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_mud_property_ref set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where mud_property_ref = :id"

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

func DeleteRMudPropertyRef(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_mud_property_ref dto.R_mud_property_ref
	r_mud_property_ref.Mud_property_ref = id

	stmt, err := db.Prepare("delete from r_mud_property_ref where mud_property_ref = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_mud_property_ref.Mud_property_ref)
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


