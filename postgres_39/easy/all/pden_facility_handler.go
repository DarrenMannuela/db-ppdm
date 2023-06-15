package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenFacility(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_facility")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_facility

	for rows.Next() {
		var pden_facility dto.Pden_facility
		if err := rows.Scan(&pden_facility.Pden_subtype, &pden_facility.Pden_id, &pden_facility.Pden_source, &pden_facility.Active_ind, &pden_facility.Effective_date, &pden_facility.Expiry_date, &pden_facility.Facility_id, &pden_facility.Facility_type, &pden_facility.No_of_gas_wells, &pden_facility.No_of_injection_wells, &pden_facility.No_of_oil_wells, &pden_facility.Ppdm_guid, &pden_facility.Remark, &pden_facility.Row_changed_by, &pden_facility.Row_changed_date, &pden_facility.Row_created_by, &pden_facility.Row_created_date, &pden_facility.Row_effective_date, &pden_facility.Row_expiry_date, &pden_facility.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_facility to result
		result = append(result, pden_facility)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenFacility(c *fiber.Ctx) error {
	var pden_facility dto.Pden_facility

	setDefaults(&pden_facility)

	if err := json.Unmarshal(c.Body(), &pden_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_facility values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	pden_facility.Row_created_date = formatDate(pden_facility.Row_created_date)
	pden_facility.Row_changed_date = formatDate(pden_facility.Row_changed_date)
	pden_facility.Effective_date = formatDateString(pden_facility.Effective_date)
	pden_facility.Expiry_date = formatDateString(pden_facility.Expiry_date)
	pden_facility.Row_effective_date = formatDateString(pden_facility.Row_effective_date)
	pden_facility.Row_expiry_date = formatDateString(pden_facility.Row_expiry_date)






	rows, err := stmt.Exec(pden_facility.Pden_subtype, pden_facility.Pden_id, pden_facility.Pden_source, pden_facility.Active_ind, pden_facility.Effective_date, pden_facility.Expiry_date, pden_facility.Facility_id, pden_facility.Facility_type, pden_facility.No_of_gas_wells, pden_facility.No_of_injection_wells, pden_facility.No_of_oil_wells, pden_facility.Ppdm_guid, pden_facility.Remark, pden_facility.Row_changed_by, pden_facility.Row_changed_date, pden_facility.Row_created_by, pden_facility.Row_created_date, pden_facility.Row_effective_date, pden_facility.Row_expiry_date, pden_facility.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenFacility(c *fiber.Ctx) error {
	var pden_facility dto.Pden_facility

	setDefaults(&pden_facility)

	if err := json.Unmarshal(c.Body(), &pden_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_facility.Pden_subtype = id

    if pden_facility.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_facility set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, facility_id = :6, facility_type = :7, no_of_gas_wells = :8, no_of_injection_wells = :9, no_of_oil_wells = :10, ppdm_guid = :11, remark = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where pden_subtype = :20")
	if err != nil {
		return err
	}

	pden_facility.Row_changed_date = formatDate(pden_facility.Row_changed_date)
	pden_facility.Effective_date = formatDateString(pden_facility.Effective_date)
	pden_facility.Expiry_date = formatDateString(pden_facility.Expiry_date)
	pden_facility.Row_effective_date = formatDateString(pden_facility.Row_effective_date)
	pden_facility.Row_expiry_date = formatDateString(pden_facility.Row_expiry_date)






	rows, err := stmt.Exec(pden_facility.Pden_id, pden_facility.Pden_source, pden_facility.Active_ind, pden_facility.Effective_date, pden_facility.Expiry_date, pden_facility.Facility_id, pden_facility.Facility_type, pden_facility.No_of_gas_wells, pden_facility.No_of_injection_wells, pden_facility.No_of_oil_wells, pden_facility.Ppdm_guid, pden_facility.Remark, pden_facility.Row_changed_by, pden_facility.Row_changed_date, pden_facility.Row_created_by, pden_facility.Row_effective_date, pden_facility.Row_expiry_date, pden_facility.Row_quality, pden_facility.Pden_subtype)
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

func PatchPdenFacility(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_facility set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_facility dto.Pden_facility
	pden_facility.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_facility where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_facility.Pden_subtype)
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


