package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_lic_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_lic_remark

	for rows.Next() {
		var facility_lic_remark dto.Facility_lic_remark
		if err := rows.Scan(&facility_lic_remark.Facility_id, &facility_lic_remark.Facility_type, &facility_lic_remark.License_id, &facility_lic_remark.Remark_id, &facility_lic_remark.Remark_seq_no, &facility_lic_remark.Active_ind, &facility_lic_remark.Author, &facility_lic_remark.Effective_date, &facility_lic_remark.Expiry_date, &facility_lic_remark.Ppdm_guid, &facility_lic_remark.Remark, &facility_lic_remark.Remark_date, &facility_lic_remark.Remark_type, &facility_lic_remark.Source, &facility_lic_remark.Row_changed_by, &facility_lic_remark.Row_changed_date, &facility_lic_remark.Row_created_by, &facility_lic_remark.Row_created_date, &facility_lic_remark.Row_effective_date, &facility_lic_remark.Row_expiry_date, &facility_lic_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_lic_remark to result
		result = append(result, facility_lic_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicRemark(c *fiber.Ctx) error {
	var facility_lic_remark dto.Facility_lic_remark

	setDefaults(&facility_lic_remark)

	if err := json.Unmarshal(c.Body(), &facility_lic_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_lic_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	facility_lic_remark.Row_created_date = formatDate(facility_lic_remark.Row_created_date)
	facility_lic_remark.Row_changed_date = formatDate(facility_lic_remark.Row_changed_date)
	facility_lic_remark.Effective_date = formatDateString(facility_lic_remark.Effective_date)
	facility_lic_remark.Expiry_date = formatDateString(facility_lic_remark.Expiry_date)
	facility_lic_remark.Remark_date = formatDateString(facility_lic_remark.Remark_date)
	facility_lic_remark.Row_effective_date = formatDateString(facility_lic_remark.Row_effective_date)
	facility_lic_remark.Row_expiry_date = formatDateString(facility_lic_remark.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_remark.Facility_id, facility_lic_remark.Facility_type, facility_lic_remark.License_id, facility_lic_remark.Remark_id, facility_lic_remark.Remark_seq_no, facility_lic_remark.Active_ind, facility_lic_remark.Author, facility_lic_remark.Effective_date, facility_lic_remark.Expiry_date, facility_lic_remark.Ppdm_guid, facility_lic_remark.Remark, facility_lic_remark.Remark_date, facility_lic_remark.Remark_type, facility_lic_remark.Source, facility_lic_remark.Row_changed_by, facility_lic_remark.Row_changed_date, facility_lic_remark.Row_created_by, facility_lic_remark.Row_created_date, facility_lic_remark.Row_effective_date, facility_lic_remark.Row_expiry_date, facility_lic_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicRemark(c *fiber.Ctx) error {
	var facility_lic_remark dto.Facility_lic_remark

	setDefaults(&facility_lic_remark)

	if err := json.Unmarshal(c.Body(), &facility_lic_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_lic_remark.Facility_id = id

    if facility_lic_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_lic_remark set facility_type = :1, license_id = :2, remark_id = :3, remark_seq_no = :4, active_ind = :5, author = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, remark_date = :11, remark_type = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where facility_id = :21")
	if err != nil {
		return err
	}

	facility_lic_remark.Row_changed_date = formatDate(facility_lic_remark.Row_changed_date)
	facility_lic_remark.Effective_date = formatDateString(facility_lic_remark.Effective_date)
	facility_lic_remark.Expiry_date = formatDateString(facility_lic_remark.Expiry_date)
	facility_lic_remark.Remark_date = formatDateString(facility_lic_remark.Remark_date)
	facility_lic_remark.Row_effective_date = formatDateString(facility_lic_remark.Row_effective_date)
	facility_lic_remark.Row_expiry_date = formatDateString(facility_lic_remark.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_remark.Facility_type, facility_lic_remark.License_id, facility_lic_remark.Remark_id, facility_lic_remark.Remark_seq_no, facility_lic_remark.Active_ind, facility_lic_remark.Author, facility_lic_remark.Effective_date, facility_lic_remark.Expiry_date, facility_lic_remark.Ppdm_guid, facility_lic_remark.Remark, facility_lic_remark.Remark_date, facility_lic_remark.Remark_type, facility_lic_remark.Source, facility_lic_remark.Row_changed_by, facility_lic_remark.Row_changed_date, facility_lic_remark.Row_created_by, facility_lic_remark.Row_effective_date, facility_lic_remark.Row_expiry_date, facility_lic_remark.Row_quality, facility_lic_remark.Facility_id)
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

func PatchFacilityLicRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_lic_remark set "
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

func DeleteFacilityLicRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_lic_remark dto.Facility_lic_remark
	facility_lic_remark.Facility_id = id

	stmt, err := db.Prepare("delete from facility_lic_remark where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_lic_remark.Facility_id)
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


