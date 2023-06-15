package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRWellStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_well_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_well_status

	for rows.Next() {
		var r_well_status dto.R_well_status
		if err := rows.Scan(&r_well_status.Status_type, &r_well_status.Status, &r_well_status.Abbreviation, &r_well_status.Active_ind, &r_well_status.Effective_date, &r_well_status.Expiry_date, &r_well_status.Long_name, &r_well_status.Ppdm_guid, &r_well_status.Remark, &r_well_status.Short_name, &r_well_status.Source, &r_well_status.Status_group, &r_well_status.Row_changed_by, &r_well_status.Row_changed_date, &r_well_status.Row_created_by, &r_well_status.Row_created_date, &r_well_status.Row_effective_date, &r_well_status.Row_expiry_date, &r_well_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_well_status to result
		result = append(result, r_well_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRWellStatus(c *fiber.Ctx) error {
	var r_well_status dto.R_well_status

	setDefaults(&r_well_status)

	if err := json.Unmarshal(c.Body(), &r_well_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_well_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	r_well_status.Row_created_date = formatDate(r_well_status.Row_created_date)
	r_well_status.Row_changed_date = formatDate(r_well_status.Row_changed_date)
	r_well_status.Effective_date = formatDateString(r_well_status.Effective_date)
	r_well_status.Expiry_date = formatDateString(r_well_status.Expiry_date)
	r_well_status.Row_effective_date = formatDateString(r_well_status.Row_effective_date)
	r_well_status.Row_expiry_date = formatDateString(r_well_status.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status.Status_type, r_well_status.Status, r_well_status.Abbreviation, r_well_status.Active_ind, r_well_status.Effective_date, r_well_status.Expiry_date, r_well_status.Long_name, r_well_status.Ppdm_guid, r_well_status.Remark, r_well_status.Short_name, r_well_status.Source, r_well_status.Status_group, r_well_status.Row_changed_by, r_well_status.Row_changed_date, r_well_status.Row_created_by, r_well_status.Row_created_date, r_well_status.Row_effective_date, r_well_status.Row_expiry_date, r_well_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRWellStatus(c *fiber.Ctx) error {
	var r_well_status dto.R_well_status

	setDefaults(&r_well_status)

	if err := json.Unmarshal(c.Body(), &r_well_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_well_status.Status_type = id

    if r_well_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_well_status set status = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, status_group = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where status_type = :19")
	if err != nil {
		return err
	}

	r_well_status.Row_changed_date = formatDate(r_well_status.Row_changed_date)
	r_well_status.Effective_date = formatDateString(r_well_status.Effective_date)
	r_well_status.Expiry_date = formatDateString(r_well_status.Expiry_date)
	r_well_status.Row_effective_date = formatDateString(r_well_status.Row_effective_date)
	r_well_status.Row_expiry_date = formatDateString(r_well_status.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status.Status, r_well_status.Abbreviation, r_well_status.Active_ind, r_well_status.Effective_date, r_well_status.Expiry_date, r_well_status.Long_name, r_well_status.Ppdm_guid, r_well_status.Remark, r_well_status.Short_name, r_well_status.Source, r_well_status.Status_group, r_well_status.Row_changed_by, r_well_status.Row_changed_date, r_well_status.Row_created_by, r_well_status.Row_effective_date, r_well_status.Row_expiry_date, r_well_status.Row_quality, r_well_status.Status_type)
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

func PatchRWellStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_well_status set "
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
	query += " where status_type = :id"

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

func DeleteRWellStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_well_status dto.R_well_status
	r_well_status.Status_type = id

	stmt, err := db.Prepare("delete from r_well_status where status_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_well_status.Status_type)
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


