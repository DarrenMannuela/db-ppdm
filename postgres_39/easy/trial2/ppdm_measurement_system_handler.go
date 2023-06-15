package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMeasurementSystem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_measurement_system")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_measurement_system

	for rows.Next() {
		var ppdm_measurement_system dto.Ppdm_measurement_system
		if err := rows.Scan(&ppdm_measurement_system.Uom_system_id, &ppdm_measurement_system.Active_ind, &ppdm_measurement_system.Effective_date, &ppdm_measurement_system.Expiry_date, &ppdm_measurement_system.Ppdm_guid, &ppdm_measurement_system.Remark, &ppdm_measurement_system.Source, &ppdm_measurement_system.Source_document_id, &ppdm_measurement_system.Unit_system_name, &ppdm_measurement_system.Row_changed_by, &ppdm_measurement_system.Row_changed_date, &ppdm_measurement_system.Row_created_by, &ppdm_measurement_system.Row_created_date, &ppdm_measurement_system.Row_effective_date, &ppdm_measurement_system.Row_expiry_date, &ppdm_measurement_system.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_measurement_system to result
		result = append(result, ppdm_measurement_system)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMeasurementSystem(c *fiber.Ctx) error {
	var ppdm_measurement_system dto.Ppdm_measurement_system

	setDefaults(&ppdm_measurement_system)

	if err := json.Unmarshal(c.Body(), &ppdm_measurement_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_measurement_system values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	ppdm_measurement_system.Row_created_date = formatDate(ppdm_measurement_system.Row_created_date)
	ppdm_measurement_system.Row_changed_date = formatDate(ppdm_measurement_system.Row_changed_date)
	ppdm_measurement_system.Effective_date = formatDateString(ppdm_measurement_system.Effective_date)
	ppdm_measurement_system.Expiry_date = formatDateString(ppdm_measurement_system.Expiry_date)
	ppdm_measurement_system.Row_effective_date = formatDateString(ppdm_measurement_system.Row_effective_date)
	ppdm_measurement_system.Row_expiry_date = formatDateString(ppdm_measurement_system.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_measurement_system.Uom_system_id, ppdm_measurement_system.Active_ind, ppdm_measurement_system.Effective_date, ppdm_measurement_system.Expiry_date, ppdm_measurement_system.Ppdm_guid, ppdm_measurement_system.Remark, ppdm_measurement_system.Source, ppdm_measurement_system.Source_document_id, ppdm_measurement_system.Unit_system_name, ppdm_measurement_system.Row_changed_by, ppdm_measurement_system.Row_changed_date, ppdm_measurement_system.Row_created_by, ppdm_measurement_system.Row_created_date, ppdm_measurement_system.Row_effective_date, ppdm_measurement_system.Row_expiry_date, ppdm_measurement_system.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMeasurementSystem(c *fiber.Ctx) error {
	var ppdm_measurement_system dto.Ppdm_measurement_system

	setDefaults(&ppdm_measurement_system)

	if err := json.Unmarshal(c.Body(), &ppdm_measurement_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_measurement_system.Uom_system_id = id

    if ppdm_measurement_system.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_measurement_system set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, remark = :5, source = :6, source_document_id = :7, unit_system_name = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where uom_system_id = :16")
	if err != nil {
		return err
	}

	ppdm_measurement_system.Row_changed_date = formatDate(ppdm_measurement_system.Row_changed_date)
	ppdm_measurement_system.Effective_date = formatDateString(ppdm_measurement_system.Effective_date)
	ppdm_measurement_system.Expiry_date = formatDateString(ppdm_measurement_system.Expiry_date)
	ppdm_measurement_system.Row_effective_date = formatDateString(ppdm_measurement_system.Row_effective_date)
	ppdm_measurement_system.Row_expiry_date = formatDateString(ppdm_measurement_system.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_measurement_system.Active_ind, ppdm_measurement_system.Effective_date, ppdm_measurement_system.Expiry_date, ppdm_measurement_system.Ppdm_guid, ppdm_measurement_system.Remark, ppdm_measurement_system.Source, ppdm_measurement_system.Source_document_id, ppdm_measurement_system.Unit_system_name, ppdm_measurement_system.Row_changed_by, ppdm_measurement_system.Row_changed_date, ppdm_measurement_system.Row_created_by, ppdm_measurement_system.Row_effective_date, ppdm_measurement_system.Row_expiry_date, ppdm_measurement_system.Row_quality, ppdm_measurement_system.Uom_system_id)
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

func PatchPpdmMeasurementSystem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_measurement_system set "
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
	query += " where uom_system_id = :id"

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

func DeletePpdmMeasurementSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_measurement_system dto.Ppdm_measurement_system
	ppdm_measurement_system.Uom_system_id = id

	stmt, err := db.Prepare("delete from ppdm_measurement_system where uom_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_measurement_system.Uom_system_id)
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


