package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenResentClass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_resent_class")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_resent_class

	for rows.Next() {
		var pden_resent_class dto.Pden_resent_class
		if err := rows.Scan(&pden_resent_class.Pden_subtype, &pden_resent_class.Pden_id, &pden_resent_class.Pden_source, &pden_resent_class.Active_ind, &pden_resent_class.Effective_date, &pden_resent_class.Expiry_date, &pden_resent_class.Long_name, &pden_resent_class.No_of_gas_wells, &pden_resent_class.No_of_injection_wells, &pden_resent_class.No_of_oil_wells, &pden_resent_class.Ppdm_guid, &pden_resent_class.Remark, &pden_resent_class.Resent_id, &pden_resent_class.Reserve_class_id, &pden_resent_class.Short_name, &pden_resent_class.Row_changed_by, &pden_resent_class.Row_changed_date, &pden_resent_class.Row_created_by, &pden_resent_class.Row_created_date, &pden_resent_class.Row_effective_date, &pden_resent_class.Row_expiry_date, &pden_resent_class.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_resent_class to result
		result = append(result, pden_resent_class)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenResentClass(c *fiber.Ctx) error {
	var pden_resent_class dto.Pden_resent_class

	setDefaults(&pden_resent_class)

	if err := json.Unmarshal(c.Body(), &pden_resent_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_resent_class values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	pden_resent_class.Row_created_date = formatDate(pden_resent_class.Row_created_date)
	pden_resent_class.Row_changed_date = formatDate(pden_resent_class.Row_changed_date)
	pden_resent_class.Effective_date = formatDateString(pden_resent_class.Effective_date)
	pden_resent_class.Expiry_date = formatDateString(pden_resent_class.Expiry_date)
	pden_resent_class.Row_effective_date = formatDateString(pden_resent_class.Row_effective_date)
	pden_resent_class.Row_expiry_date = formatDateString(pden_resent_class.Row_expiry_date)






	rows, err := stmt.Exec(pden_resent_class.Pden_subtype, pden_resent_class.Pden_id, pden_resent_class.Pden_source, pden_resent_class.Active_ind, pden_resent_class.Effective_date, pden_resent_class.Expiry_date, pden_resent_class.Long_name, pden_resent_class.No_of_gas_wells, pden_resent_class.No_of_injection_wells, pden_resent_class.No_of_oil_wells, pden_resent_class.Ppdm_guid, pden_resent_class.Remark, pden_resent_class.Resent_id, pden_resent_class.Reserve_class_id, pden_resent_class.Short_name, pden_resent_class.Row_changed_by, pden_resent_class.Row_changed_date, pden_resent_class.Row_created_by, pden_resent_class.Row_created_date, pden_resent_class.Row_effective_date, pden_resent_class.Row_expiry_date, pden_resent_class.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenResentClass(c *fiber.Ctx) error {
	var pden_resent_class dto.Pden_resent_class

	setDefaults(&pden_resent_class)

	if err := json.Unmarshal(c.Body(), &pden_resent_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_resent_class.Pden_subtype = id

    if pden_resent_class.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_resent_class set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, no_of_gas_wells = :7, no_of_injection_wells = :8, no_of_oil_wells = :9, ppdm_guid = :10, remark = :11, resent_id = :12, reserve_class_id = :13, short_name = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where pden_subtype = :22")
	if err != nil {
		return err
	}

	pden_resent_class.Row_changed_date = formatDate(pden_resent_class.Row_changed_date)
	pden_resent_class.Effective_date = formatDateString(pden_resent_class.Effective_date)
	pden_resent_class.Expiry_date = formatDateString(pden_resent_class.Expiry_date)
	pden_resent_class.Row_effective_date = formatDateString(pden_resent_class.Row_effective_date)
	pden_resent_class.Row_expiry_date = formatDateString(pden_resent_class.Row_expiry_date)






	rows, err := stmt.Exec(pden_resent_class.Pden_id, pden_resent_class.Pden_source, pden_resent_class.Active_ind, pden_resent_class.Effective_date, pden_resent_class.Expiry_date, pden_resent_class.Long_name, pden_resent_class.No_of_gas_wells, pden_resent_class.No_of_injection_wells, pden_resent_class.No_of_oil_wells, pden_resent_class.Ppdm_guid, pden_resent_class.Remark, pden_resent_class.Resent_id, pden_resent_class.Reserve_class_id, pden_resent_class.Short_name, pden_resent_class.Row_changed_by, pden_resent_class.Row_changed_date, pden_resent_class.Row_created_by, pden_resent_class.Row_effective_date, pden_resent_class.Row_expiry_date, pden_resent_class.Row_quality, pden_resent_class.Pden_subtype)
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

func PatchPdenResentClass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_resent_class set "
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

func DeletePdenResentClass(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_resent_class dto.Pden_resent_class
	pden_resent_class.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_resent_class where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_resent_class.Pden_subtype)
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


