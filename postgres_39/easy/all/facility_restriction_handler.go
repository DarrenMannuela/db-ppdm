package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityRestriction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_restriction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_restriction

	for rows.Next() {
		var facility_restriction dto.Facility_restriction
		if err := rows.Scan(&facility_restriction.Facility_id, &facility_restriction.Facility_type, &facility_restriction.Restriction_id, &facility_restriction.Restriction_version, &facility_restriction.Active_ind, &facility_restriction.Description, &facility_restriction.Effective_date, &facility_restriction.Expiry_date, &facility_restriction.Ppdm_guid, &facility_restriction.Remark, &facility_restriction.Source, &facility_restriction.Row_changed_by, &facility_restriction.Row_changed_date, &facility_restriction.Row_created_by, &facility_restriction.Row_created_date, &facility_restriction.Row_effective_date, &facility_restriction.Row_expiry_date, &facility_restriction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_restriction to result
		result = append(result, facility_restriction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityRestriction(c *fiber.Ctx) error {
	var facility_restriction dto.Facility_restriction

	setDefaults(&facility_restriction)

	if err := json.Unmarshal(c.Body(), &facility_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_restriction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	facility_restriction.Row_created_date = formatDate(facility_restriction.Row_created_date)
	facility_restriction.Row_changed_date = formatDate(facility_restriction.Row_changed_date)
	facility_restriction.Effective_date = formatDateString(facility_restriction.Effective_date)
	facility_restriction.Expiry_date = formatDateString(facility_restriction.Expiry_date)
	facility_restriction.Row_effective_date = formatDateString(facility_restriction.Row_effective_date)
	facility_restriction.Row_expiry_date = formatDateString(facility_restriction.Row_expiry_date)






	rows, err := stmt.Exec(facility_restriction.Facility_id, facility_restriction.Facility_type, facility_restriction.Restriction_id, facility_restriction.Restriction_version, facility_restriction.Active_ind, facility_restriction.Description, facility_restriction.Effective_date, facility_restriction.Expiry_date, facility_restriction.Ppdm_guid, facility_restriction.Remark, facility_restriction.Source, facility_restriction.Row_changed_by, facility_restriction.Row_changed_date, facility_restriction.Row_created_by, facility_restriction.Row_created_date, facility_restriction.Row_effective_date, facility_restriction.Row_expiry_date, facility_restriction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityRestriction(c *fiber.Ctx) error {
	var facility_restriction dto.Facility_restriction

	setDefaults(&facility_restriction)

	if err := json.Unmarshal(c.Body(), &facility_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_restriction.Facility_id = id

    if facility_restriction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_restriction set facility_type = :1, restriction_id = :2, restriction_version = :3, active_ind = :4, description = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where facility_id = :18")
	if err != nil {
		return err
	}

	facility_restriction.Row_changed_date = formatDate(facility_restriction.Row_changed_date)
	facility_restriction.Effective_date = formatDateString(facility_restriction.Effective_date)
	facility_restriction.Expiry_date = formatDateString(facility_restriction.Expiry_date)
	facility_restriction.Row_effective_date = formatDateString(facility_restriction.Row_effective_date)
	facility_restriction.Row_expiry_date = formatDateString(facility_restriction.Row_expiry_date)






	rows, err := stmt.Exec(facility_restriction.Facility_type, facility_restriction.Restriction_id, facility_restriction.Restriction_version, facility_restriction.Active_ind, facility_restriction.Description, facility_restriction.Effective_date, facility_restriction.Expiry_date, facility_restriction.Ppdm_guid, facility_restriction.Remark, facility_restriction.Source, facility_restriction.Row_changed_by, facility_restriction.Row_changed_date, facility_restriction.Row_created_by, facility_restriction.Row_effective_date, facility_restriction.Row_expiry_date, facility_restriction.Row_quality, facility_restriction.Facility_id)
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

func PatchFacilityRestriction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_restriction set "
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

func DeleteFacilityRestriction(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_restriction dto.Facility_restriction
	facility_restriction.Facility_id = id

	stmt, err := db.Prepare("delete from facility_restriction where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_restriction.Facility_id)
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


