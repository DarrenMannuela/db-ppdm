package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSwApplication(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_sw_application")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_sw_application

	for rows.Next() {
		var ppdm_sw_application dto.Ppdm_sw_application
		if err := rows.Scan(&ppdm_sw_application.Sw_application_id, &ppdm_sw_application.Abbreviation, &ppdm_sw_application.Active_ind, &ppdm_sw_application.Application_version, &ppdm_sw_application.Effective_date, &ppdm_sw_application.Expiry_date, &ppdm_sw_application.Function_description, &ppdm_sw_application.Long_name, &ppdm_sw_application.Owner_ba_id, &ppdm_sw_application.Ppdm_guid, &ppdm_sw_application.Preferred_uom_system_id, &ppdm_sw_application.Remark, &ppdm_sw_application.Short_name, &ppdm_sw_application.Source, &ppdm_sw_application.Row_changed_by, &ppdm_sw_application.Row_changed_date, &ppdm_sw_application.Row_created_by, &ppdm_sw_application.Row_created_date, &ppdm_sw_application.Row_effective_date, &ppdm_sw_application.Row_expiry_date, &ppdm_sw_application.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_sw_application to result
		result = append(result, ppdm_sw_application)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSwApplication(c *fiber.Ctx) error {
	var ppdm_sw_application dto.Ppdm_sw_application

	setDefaults(&ppdm_sw_application)

	if err := json.Unmarshal(c.Body(), &ppdm_sw_application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_sw_application values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ppdm_sw_application.Row_created_date = formatDate(ppdm_sw_application.Row_created_date)
	ppdm_sw_application.Row_changed_date = formatDate(ppdm_sw_application.Row_changed_date)
	ppdm_sw_application.Effective_date = formatDateString(ppdm_sw_application.Effective_date)
	ppdm_sw_application.Expiry_date = formatDateString(ppdm_sw_application.Expiry_date)
	ppdm_sw_application.Row_effective_date = formatDateString(ppdm_sw_application.Row_effective_date)
	ppdm_sw_application.Row_expiry_date = formatDateString(ppdm_sw_application.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_sw_application.Sw_application_id, ppdm_sw_application.Abbreviation, ppdm_sw_application.Active_ind, ppdm_sw_application.Application_version, ppdm_sw_application.Effective_date, ppdm_sw_application.Expiry_date, ppdm_sw_application.Function_description, ppdm_sw_application.Long_name, ppdm_sw_application.Owner_ba_id, ppdm_sw_application.Ppdm_guid, ppdm_sw_application.Preferred_uom_system_id, ppdm_sw_application.Remark, ppdm_sw_application.Short_name, ppdm_sw_application.Source, ppdm_sw_application.Row_changed_by, ppdm_sw_application.Row_changed_date, ppdm_sw_application.Row_created_by, ppdm_sw_application.Row_created_date, ppdm_sw_application.Row_effective_date, ppdm_sw_application.Row_expiry_date, ppdm_sw_application.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSwApplication(c *fiber.Ctx) error {
	var ppdm_sw_application dto.Ppdm_sw_application

	setDefaults(&ppdm_sw_application)

	if err := json.Unmarshal(c.Body(), &ppdm_sw_application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_sw_application.Sw_application_id = id

    if ppdm_sw_application.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_sw_application set abbreviation = :1, active_ind = :2, application_version = :3, effective_date = :4, expiry_date = :5, function_description = :6, long_name = :7, owner_ba_id = :8, ppdm_guid = :9, preferred_uom_system_id = :10, remark = :11, short_name = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where sw_application_id = :21")
	if err != nil {
		return err
	}

	ppdm_sw_application.Row_changed_date = formatDate(ppdm_sw_application.Row_changed_date)
	ppdm_sw_application.Effective_date = formatDateString(ppdm_sw_application.Effective_date)
	ppdm_sw_application.Expiry_date = formatDateString(ppdm_sw_application.Expiry_date)
	ppdm_sw_application.Row_effective_date = formatDateString(ppdm_sw_application.Row_effective_date)
	ppdm_sw_application.Row_expiry_date = formatDateString(ppdm_sw_application.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_sw_application.Abbreviation, ppdm_sw_application.Active_ind, ppdm_sw_application.Application_version, ppdm_sw_application.Effective_date, ppdm_sw_application.Expiry_date, ppdm_sw_application.Function_description, ppdm_sw_application.Long_name, ppdm_sw_application.Owner_ba_id, ppdm_sw_application.Ppdm_guid, ppdm_sw_application.Preferred_uom_system_id, ppdm_sw_application.Remark, ppdm_sw_application.Short_name, ppdm_sw_application.Source, ppdm_sw_application.Row_changed_by, ppdm_sw_application.Row_changed_date, ppdm_sw_application.Row_created_by, ppdm_sw_application.Row_effective_date, ppdm_sw_application.Row_expiry_date, ppdm_sw_application.Row_quality, ppdm_sw_application.Sw_application_id)
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

func PatchPpdmSwApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_sw_application set "
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
	query += " where sw_application_id = :id"

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

func DeletePpdmSwApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_sw_application dto.Ppdm_sw_application
	ppdm_sw_application.Sw_application_id = id

	stmt, err := db.Prepare("delete from ppdm_sw_application where sw_application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_sw_application.Sw_application_id)
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


