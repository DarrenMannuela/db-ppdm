package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenBusinessAssoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_business_assoc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_business_assoc

	for rows.Next() {
		var pden_business_assoc dto.Pden_business_assoc
		if err := rows.Scan(&pden_business_assoc.Pden_subtype, &pden_business_assoc.Pden_id, &pden_business_assoc.Pden_source, &pden_business_assoc.Active_ind, &pden_business_assoc.Business_associate_id, &pden_business_assoc.Effective_date, &pden_business_assoc.Expiry_date, &pden_business_assoc.No_of_gas_wells, &pden_business_assoc.No_of_injection_wells, &pden_business_assoc.No_of_oil_wells, &pden_business_assoc.Ppdm_guid, &pden_business_assoc.Remark, &pden_business_assoc.Row_changed_by, &pden_business_assoc.Row_changed_date, &pden_business_assoc.Row_created_by, &pden_business_assoc.Row_created_date, &pden_business_assoc.Row_effective_date, &pden_business_assoc.Row_expiry_date, &pden_business_assoc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_business_assoc to result
		result = append(result, pden_business_assoc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenBusinessAssoc(c *fiber.Ctx) error {
	var pden_business_assoc dto.Pden_business_assoc

	setDefaults(&pden_business_assoc)

	if err := json.Unmarshal(c.Body(), &pden_business_assoc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_business_assoc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	pden_business_assoc.Row_created_date = formatDate(pden_business_assoc.Row_created_date)
	pden_business_assoc.Row_changed_date = formatDate(pden_business_assoc.Row_changed_date)
	pden_business_assoc.Effective_date = formatDateString(pden_business_assoc.Effective_date)
	pden_business_assoc.Expiry_date = formatDateString(pden_business_assoc.Expiry_date)
	pden_business_assoc.Row_effective_date = formatDateString(pden_business_assoc.Row_effective_date)
	pden_business_assoc.Row_expiry_date = formatDateString(pden_business_assoc.Row_expiry_date)






	rows, err := stmt.Exec(pden_business_assoc.Pden_subtype, pden_business_assoc.Pden_id, pden_business_assoc.Pden_source, pden_business_assoc.Active_ind, pden_business_assoc.Business_associate_id, pden_business_assoc.Effective_date, pden_business_assoc.Expiry_date, pden_business_assoc.No_of_gas_wells, pden_business_assoc.No_of_injection_wells, pden_business_assoc.No_of_oil_wells, pden_business_assoc.Ppdm_guid, pden_business_assoc.Remark, pden_business_assoc.Row_changed_by, pden_business_assoc.Row_changed_date, pden_business_assoc.Row_created_by, pden_business_assoc.Row_created_date, pden_business_assoc.Row_effective_date, pden_business_assoc.Row_expiry_date, pden_business_assoc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenBusinessAssoc(c *fiber.Ctx) error {
	var pden_business_assoc dto.Pden_business_assoc

	setDefaults(&pden_business_assoc)

	if err := json.Unmarshal(c.Body(), &pden_business_assoc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_business_assoc.Pden_subtype = id

    if pden_business_assoc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_business_assoc set pden_id = :1, pden_source = :2, active_ind = :3, business_associate_id = :4, effective_date = :5, expiry_date = :6, no_of_gas_wells = :7, no_of_injection_wells = :8, no_of_oil_wells = :9, ppdm_guid = :10, remark = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where pden_subtype = :19")
	if err != nil {
		return err
	}

	pden_business_assoc.Row_changed_date = formatDate(pden_business_assoc.Row_changed_date)
	pden_business_assoc.Effective_date = formatDateString(pden_business_assoc.Effective_date)
	pden_business_assoc.Expiry_date = formatDateString(pden_business_assoc.Expiry_date)
	pden_business_assoc.Row_effective_date = formatDateString(pden_business_assoc.Row_effective_date)
	pden_business_assoc.Row_expiry_date = formatDateString(pden_business_assoc.Row_expiry_date)






	rows, err := stmt.Exec(pden_business_assoc.Pden_id, pden_business_assoc.Pden_source, pden_business_assoc.Active_ind, pden_business_assoc.Business_associate_id, pden_business_assoc.Effective_date, pden_business_assoc.Expiry_date, pden_business_assoc.No_of_gas_wells, pden_business_assoc.No_of_injection_wells, pden_business_assoc.No_of_oil_wells, pden_business_assoc.Ppdm_guid, pden_business_assoc.Remark, pden_business_assoc.Row_changed_by, pden_business_assoc.Row_changed_date, pden_business_assoc.Row_created_by, pden_business_assoc.Row_effective_date, pden_business_assoc.Row_expiry_date, pden_business_assoc.Row_quality, pden_business_assoc.Pden_subtype)
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

func PatchPdenBusinessAssoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_business_assoc set "
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

func DeletePdenBusinessAssoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_business_assoc dto.Pden_business_assoc
	pden_business_assoc.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_business_assoc where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_business_assoc.Pden_subtype)
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


