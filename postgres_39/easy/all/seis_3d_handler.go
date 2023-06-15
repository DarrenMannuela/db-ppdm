package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeis3D(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_3d")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_3d

	for rows.Next() {
		var seis_3d dto.Seis_3d
		if err := rows.Scan(&seis_3d.Seis_set_subtype, &seis_3d.Seis_3d_id, &seis_3d.Active_ind, &seis_3d.Effective_date, &seis_3d.Expiry_date, &seis_3d.Ppdm_guid, &seis_3d.Remark, &seis_3d.Seis_3d_reason, &seis_3d.Seis_3d_type, &seis_3d.Source, &seis_3d.Row_changed_by, &seis_3d.Row_changed_date, &seis_3d.Row_created_by, &seis_3d.Row_created_date, &seis_3d.Row_effective_date, &seis_3d.Row_expiry_date, &seis_3d.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_3d to result
		result = append(result, seis_3d)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeis3D(c *fiber.Ctx) error {
	var seis_3d dto.Seis_3d

	setDefaults(&seis_3d)

	if err := json.Unmarshal(c.Body(), &seis_3d); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_3d values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	seis_3d.Row_created_date = formatDate(seis_3d.Row_created_date)
	seis_3d.Row_changed_date = formatDate(seis_3d.Row_changed_date)
	seis_3d.Effective_date = formatDateString(seis_3d.Effective_date)
	seis_3d.Expiry_date = formatDateString(seis_3d.Expiry_date)
	seis_3d.Row_effective_date = formatDateString(seis_3d.Row_effective_date)
	seis_3d.Row_expiry_date = formatDateString(seis_3d.Row_expiry_date)






	rows, err := stmt.Exec(seis_3d.Seis_set_subtype, seis_3d.Seis_3d_id, seis_3d.Active_ind, seis_3d.Effective_date, seis_3d.Expiry_date, seis_3d.Ppdm_guid, seis_3d.Remark, seis_3d.Seis_3d_reason, seis_3d.Seis_3d_type, seis_3d.Source, seis_3d.Row_changed_by, seis_3d.Row_changed_date, seis_3d.Row_created_by, seis_3d.Row_created_date, seis_3d.Row_effective_date, seis_3d.Row_expiry_date, seis_3d.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeis3D(c *fiber.Ctx) error {
	var seis_3d dto.Seis_3d

	setDefaults(&seis_3d)

	if err := json.Unmarshal(c.Body(), &seis_3d); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_3d.Seis_set_subtype = id

    if seis_3d.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_3d set seis_3d_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, seis_3d_reason = :7, seis_3d_type = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where seis_set_subtype = :17")
	if err != nil {
		return err
	}

	seis_3d.Row_changed_date = formatDate(seis_3d.Row_changed_date)
	seis_3d.Effective_date = formatDateString(seis_3d.Effective_date)
	seis_3d.Expiry_date = formatDateString(seis_3d.Expiry_date)
	seis_3d.Row_effective_date = formatDateString(seis_3d.Row_effective_date)
	seis_3d.Row_expiry_date = formatDateString(seis_3d.Row_expiry_date)






	rows, err := stmt.Exec(seis_3d.Seis_3d_id, seis_3d.Active_ind, seis_3d.Effective_date, seis_3d.Expiry_date, seis_3d.Ppdm_guid, seis_3d.Remark, seis_3d.Seis_3d_reason, seis_3d.Seis_3d_type, seis_3d.Source, seis_3d.Row_changed_by, seis_3d.Row_changed_date, seis_3d.Row_created_by, seis_3d.Row_effective_date, seis_3d.Row_expiry_date, seis_3d.Row_quality, seis_3d.Seis_set_subtype)
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

func PatchSeis3D(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_3d set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeis3D(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_3d dto.Seis_3d
	seis_3d.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_3d where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_3d.Seis_set_subtype)
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


