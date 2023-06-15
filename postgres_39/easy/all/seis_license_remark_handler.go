package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisLicenseRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_license_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_license_remark

	for rows.Next() {
		var seis_license_remark dto.Seis_license_remark
		if err := rows.Scan(&seis_license_remark.Seis_set_subtype, &seis_license_remark.Seis_set_id, &seis_license_remark.License_id, &seis_license_remark.Remark_id, &seis_license_remark.Remark_seq_no, &seis_license_remark.Active_ind, &seis_license_remark.Author, &seis_license_remark.Effective_date, &seis_license_remark.Expiry_date, &seis_license_remark.Ppdm_guid, &seis_license_remark.Remark, &seis_license_remark.Remark_date, &seis_license_remark.Remark_type, &seis_license_remark.Source, &seis_license_remark.Row_changed_by, &seis_license_remark.Row_changed_date, &seis_license_remark.Row_created_by, &seis_license_remark.Row_created_date, &seis_license_remark.Row_effective_date, &seis_license_remark.Row_expiry_date, &seis_license_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_license_remark to result
		result = append(result, seis_license_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisLicenseRemark(c *fiber.Ctx) error {
	var seis_license_remark dto.Seis_license_remark

	setDefaults(&seis_license_remark)

	if err := json.Unmarshal(c.Body(), &seis_license_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_license_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	seis_license_remark.Row_created_date = formatDate(seis_license_remark.Row_created_date)
	seis_license_remark.Row_changed_date = formatDate(seis_license_remark.Row_changed_date)
	seis_license_remark.Effective_date = formatDateString(seis_license_remark.Effective_date)
	seis_license_remark.Expiry_date = formatDateString(seis_license_remark.Expiry_date)
	seis_license_remark.Remark_date = formatDateString(seis_license_remark.Remark_date)
	seis_license_remark.Row_effective_date = formatDateString(seis_license_remark.Row_effective_date)
	seis_license_remark.Row_expiry_date = formatDateString(seis_license_remark.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_remark.Seis_set_subtype, seis_license_remark.Seis_set_id, seis_license_remark.License_id, seis_license_remark.Remark_id, seis_license_remark.Remark_seq_no, seis_license_remark.Active_ind, seis_license_remark.Author, seis_license_remark.Effective_date, seis_license_remark.Expiry_date, seis_license_remark.Ppdm_guid, seis_license_remark.Remark, seis_license_remark.Remark_date, seis_license_remark.Remark_type, seis_license_remark.Source, seis_license_remark.Row_changed_by, seis_license_remark.Row_changed_date, seis_license_remark.Row_created_by, seis_license_remark.Row_created_date, seis_license_remark.Row_effective_date, seis_license_remark.Row_expiry_date, seis_license_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisLicenseRemark(c *fiber.Ctx) error {
	var seis_license_remark dto.Seis_license_remark

	setDefaults(&seis_license_remark)

	if err := json.Unmarshal(c.Body(), &seis_license_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_license_remark.Seis_set_subtype = id

    if seis_license_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_license_remark set seis_set_id = :1, license_id = :2, remark_id = :3, remark_seq_no = :4, active_ind = :5, author = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, remark_date = :11, remark_type = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where seis_set_subtype = :21")
	if err != nil {
		return err
	}

	seis_license_remark.Row_changed_date = formatDate(seis_license_remark.Row_changed_date)
	seis_license_remark.Effective_date = formatDateString(seis_license_remark.Effective_date)
	seis_license_remark.Expiry_date = formatDateString(seis_license_remark.Expiry_date)
	seis_license_remark.Remark_date = formatDateString(seis_license_remark.Remark_date)
	seis_license_remark.Row_effective_date = formatDateString(seis_license_remark.Row_effective_date)
	seis_license_remark.Row_expiry_date = formatDateString(seis_license_remark.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_remark.Seis_set_id, seis_license_remark.License_id, seis_license_remark.Remark_id, seis_license_remark.Remark_seq_no, seis_license_remark.Active_ind, seis_license_remark.Author, seis_license_remark.Effective_date, seis_license_remark.Expiry_date, seis_license_remark.Ppdm_guid, seis_license_remark.Remark, seis_license_remark.Remark_date, seis_license_remark.Remark_type, seis_license_remark.Source, seis_license_remark.Row_changed_by, seis_license_remark.Row_changed_date, seis_license_remark.Row_created_by, seis_license_remark.Row_effective_date, seis_license_remark.Row_expiry_date, seis_license_remark.Row_quality, seis_license_remark.Seis_set_subtype)
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

func PatchSeisLicenseRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_license_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisLicenseRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_license_remark dto.Seis_license_remark
	seis_license_remark.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_license_remark where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_license_remark.Seis_set_subtype)
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


