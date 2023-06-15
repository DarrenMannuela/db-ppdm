package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLicenseRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_license_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_license_remark

	for rows.Next() {
		var well_license_remark dto.Well_license_remark
		if err := rows.Scan(&well_license_remark.Uwi, &well_license_remark.Well_lic_source, &well_license_remark.License_id, &well_license_remark.Remark_id, &well_license_remark.Remark_seq_no, &well_license_remark.Active_ind, &well_license_remark.Author, &well_license_remark.Effective_date, &well_license_remark.Expiry_date, &well_license_remark.Ppdm_guid, &well_license_remark.Remark, &well_license_remark.Remark_date, &well_license_remark.Remark_type, &well_license_remark.Source, &well_license_remark.Row_changed_by, &well_license_remark.Row_changed_date, &well_license_remark.Row_created_by, &well_license_remark.Row_created_date, &well_license_remark.Row_effective_date, &well_license_remark.Row_expiry_date, &well_license_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_license_remark to result
		result = append(result, well_license_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLicenseRemark(c *fiber.Ctx) error {
	var well_license_remark dto.Well_license_remark

	setDefaults(&well_license_remark)

	if err := json.Unmarshal(c.Body(), &well_license_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_license_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_license_remark.Row_created_date = formatDate(well_license_remark.Row_created_date)
	well_license_remark.Row_changed_date = formatDate(well_license_remark.Row_changed_date)
	well_license_remark.Effective_date = formatDateString(well_license_remark.Effective_date)
	well_license_remark.Expiry_date = formatDateString(well_license_remark.Expiry_date)
	well_license_remark.Remark_date = formatDateString(well_license_remark.Remark_date)
	well_license_remark.Row_effective_date = formatDateString(well_license_remark.Row_effective_date)
	well_license_remark.Row_expiry_date = formatDateString(well_license_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_license_remark.Uwi, well_license_remark.Well_lic_source, well_license_remark.License_id, well_license_remark.Remark_id, well_license_remark.Remark_seq_no, well_license_remark.Active_ind, well_license_remark.Author, well_license_remark.Effective_date, well_license_remark.Expiry_date, well_license_remark.Ppdm_guid, well_license_remark.Remark, well_license_remark.Remark_date, well_license_remark.Remark_type, well_license_remark.Source, well_license_remark.Row_changed_by, well_license_remark.Row_changed_date, well_license_remark.Row_created_by, well_license_remark.Row_created_date, well_license_remark.Row_effective_date, well_license_remark.Row_expiry_date, well_license_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLicenseRemark(c *fiber.Ctx) error {
	var well_license_remark dto.Well_license_remark

	setDefaults(&well_license_remark)

	if err := json.Unmarshal(c.Body(), &well_license_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_license_remark.Uwi = id

    if well_license_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_license_remark set well_lic_source = :1, license_id = :2, remark_id = :3, remark_seq_no = :4, active_ind = :5, author = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, remark_date = :11, remark_type = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_license_remark.Row_changed_date = formatDate(well_license_remark.Row_changed_date)
	well_license_remark.Effective_date = formatDateString(well_license_remark.Effective_date)
	well_license_remark.Expiry_date = formatDateString(well_license_remark.Expiry_date)
	well_license_remark.Remark_date = formatDateString(well_license_remark.Remark_date)
	well_license_remark.Row_effective_date = formatDateString(well_license_remark.Row_effective_date)
	well_license_remark.Row_expiry_date = formatDateString(well_license_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_license_remark.Well_lic_source, well_license_remark.License_id, well_license_remark.Remark_id, well_license_remark.Remark_seq_no, well_license_remark.Active_ind, well_license_remark.Author, well_license_remark.Effective_date, well_license_remark.Expiry_date, well_license_remark.Ppdm_guid, well_license_remark.Remark, well_license_remark.Remark_date, well_license_remark.Remark_type, well_license_remark.Source, well_license_remark.Row_changed_by, well_license_remark.Row_changed_date, well_license_remark.Row_created_by, well_license_remark.Row_effective_date, well_license_remark.Row_expiry_date, well_license_remark.Row_quality, well_license_remark.Uwi)
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

func PatchWellLicenseRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_license_remark set "
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
	query += " where uwi = :id"

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

func DeleteWellLicenseRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_license_remark dto.Well_license_remark
	well_license_remark.Uwi = id

	stmt, err := db.Prepare("delete from well_license_remark where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_license_remark.Uwi)
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


