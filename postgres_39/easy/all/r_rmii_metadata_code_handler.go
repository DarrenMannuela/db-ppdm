package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRRmiiMetadataCode(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_rmii_metadata_code")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_rmii_metadata_code

	for rows.Next() {
		var r_rmii_metadata_code dto.R_rmii_metadata_code
		if err := rows.Scan(&r_rmii_metadata_code.Metadata_type, &r_rmii_metadata_code.Metadata_code, &r_rmii_metadata_code.Abbreviation, &r_rmii_metadata_code.Active_ind, &r_rmii_metadata_code.Effective_date, &r_rmii_metadata_code.Expiry_date, &r_rmii_metadata_code.Long_name, &r_rmii_metadata_code.Ppdm_guid, &r_rmii_metadata_code.Remark, &r_rmii_metadata_code.Short_name, &r_rmii_metadata_code.Source, &r_rmii_metadata_code.Row_changed_by, &r_rmii_metadata_code.Row_changed_date, &r_rmii_metadata_code.Row_created_by, &r_rmii_metadata_code.Row_created_date, &r_rmii_metadata_code.Row_effective_date, &r_rmii_metadata_code.Row_expiry_date, &r_rmii_metadata_code.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_rmii_metadata_code to result
		result = append(result, r_rmii_metadata_code)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRRmiiMetadataCode(c *fiber.Ctx) error {
	var r_rmii_metadata_code dto.R_rmii_metadata_code

	setDefaults(&r_rmii_metadata_code)

	if err := json.Unmarshal(c.Body(), &r_rmii_metadata_code); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_rmii_metadata_code values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_rmii_metadata_code.Row_created_date = formatDate(r_rmii_metadata_code.Row_created_date)
	r_rmii_metadata_code.Row_changed_date = formatDate(r_rmii_metadata_code.Row_changed_date)
	r_rmii_metadata_code.Effective_date = formatDateString(r_rmii_metadata_code.Effective_date)
	r_rmii_metadata_code.Expiry_date = formatDateString(r_rmii_metadata_code.Expiry_date)
	r_rmii_metadata_code.Row_effective_date = formatDateString(r_rmii_metadata_code.Row_effective_date)
	r_rmii_metadata_code.Row_expiry_date = formatDateString(r_rmii_metadata_code.Row_expiry_date)






	rows, err := stmt.Exec(r_rmii_metadata_code.Metadata_type, r_rmii_metadata_code.Metadata_code, r_rmii_metadata_code.Abbreviation, r_rmii_metadata_code.Active_ind, r_rmii_metadata_code.Effective_date, r_rmii_metadata_code.Expiry_date, r_rmii_metadata_code.Long_name, r_rmii_metadata_code.Ppdm_guid, r_rmii_metadata_code.Remark, r_rmii_metadata_code.Short_name, r_rmii_metadata_code.Source, r_rmii_metadata_code.Row_changed_by, r_rmii_metadata_code.Row_changed_date, r_rmii_metadata_code.Row_created_by, r_rmii_metadata_code.Row_created_date, r_rmii_metadata_code.Row_effective_date, r_rmii_metadata_code.Row_expiry_date, r_rmii_metadata_code.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRRmiiMetadataCode(c *fiber.Ctx) error {
	var r_rmii_metadata_code dto.R_rmii_metadata_code

	setDefaults(&r_rmii_metadata_code)

	if err := json.Unmarshal(c.Body(), &r_rmii_metadata_code); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_rmii_metadata_code.Metadata_type = id

    if r_rmii_metadata_code.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_rmii_metadata_code set metadata_code = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where metadata_type = :18")
	if err != nil {
		return err
	}

	r_rmii_metadata_code.Row_changed_date = formatDate(r_rmii_metadata_code.Row_changed_date)
	r_rmii_metadata_code.Effective_date = formatDateString(r_rmii_metadata_code.Effective_date)
	r_rmii_metadata_code.Expiry_date = formatDateString(r_rmii_metadata_code.Expiry_date)
	r_rmii_metadata_code.Row_effective_date = formatDateString(r_rmii_metadata_code.Row_effective_date)
	r_rmii_metadata_code.Row_expiry_date = formatDateString(r_rmii_metadata_code.Row_expiry_date)






	rows, err := stmt.Exec(r_rmii_metadata_code.Metadata_code, r_rmii_metadata_code.Abbreviation, r_rmii_metadata_code.Active_ind, r_rmii_metadata_code.Effective_date, r_rmii_metadata_code.Expiry_date, r_rmii_metadata_code.Long_name, r_rmii_metadata_code.Ppdm_guid, r_rmii_metadata_code.Remark, r_rmii_metadata_code.Short_name, r_rmii_metadata_code.Source, r_rmii_metadata_code.Row_changed_by, r_rmii_metadata_code.Row_changed_date, r_rmii_metadata_code.Row_created_by, r_rmii_metadata_code.Row_effective_date, r_rmii_metadata_code.Row_expiry_date, r_rmii_metadata_code.Row_quality, r_rmii_metadata_code.Metadata_type)
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

func PatchRRmiiMetadataCode(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_rmii_metadata_code set "
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
	query += " where metadata_type = :id"

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

func DeleteRRmiiMetadataCode(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_rmii_metadata_code dto.R_rmii_metadata_code
	r_rmii_metadata_code.Metadata_type = id

	stmt, err := db.Prepare("delete from r_rmii_metadata_code where metadata_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_rmii_metadata_code.Metadata_type)
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

