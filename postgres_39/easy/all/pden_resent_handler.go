package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenResent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_resent")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_resent

	for rows.Next() {
		var pden_resent dto.Pden_resent
		if err := rows.Scan(&pden_resent.Pden_subtype, &pden_resent.Pden_id, &pden_resent.Pden_source, &pden_resent.Active_ind, &pden_resent.Effective_date, &pden_resent.Expiry_date, &pden_resent.Long_name, &pden_resent.No_of_gas_wells, &pden_resent.No_of_injection_wells, &pden_resent.No_of_oil_wells, &pden_resent.Ppdm_guid, &pden_resent.Remark, &pden_resent.Resent_id, &pden_resent.Short_name, &pden_resent.Row_changed_by, &pden_resent.Row_changed_date, &pden_resent.Row_created_by, &pden_resent.Row_created_date, &pden_resent.Row_effective_date, &pden_resent.Row_expiry_date, &pden_resent.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_resent to result
		result = append(result, pden_resent)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenResent(c *fiber.Ctx) error {
	var pden_resent dto.Pden_resent

	setDefaults(&pden_resent)

	if err := json.Unmarshal(c.Body(), &pden_resent); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_resent values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	pden_resent.Row_created_date = formatDate(pden_resent.Row_created_date)
	pden_resent.Row_changed_date = formatDate(pden_resent.Row_changed_date)
	pden_resent.Effective_date = formatDateString(pden_resent.Effective_date)
	pden_resent.Expiry_date = formatDateString(pden_resent.Expiry_date)
	pden_resent.Row_effective_date = formatDateString(pden_resent.Row_effective_date)
	pden_resent.Row_expiry_date = formatDateString(pden_resent.Row_expiry_date)






	rows, err := stmt.Exec(pden_resent.Pden_subtype, pden_resent.Pden_id, pden_resent.Pden_source, pden_resent.Active_ind, pden_resent.Effective_date, pden_resent.Expiry_date, pden_resent.Long_name, pden_resent.No_of_gas_wells, pden_resent.No_of_injection_wells, pden_resent.No_of_oil_wells, pden_resent.Ppdm_guid, pden_resent.Remark, pden_resent.Resent_id, pden_resent.Short_name, pden_resent.Row_changed_by, pden_resent.Row_changed_date, pden_resent.Row_created_by, pden_resent.Row_created_date, pden_resent.Row_effective_date, pden_resent.Row_expiry_date, pden_resent.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenResent(c *fiber.Ctx) error {
	var pden_resent dto.Pden_resent

	setDefaults(&pden_resent)

	if err := json.Unmarshal(c.Body(), &pden_resent); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_resent.Pden_subtype = id

    if pden_resent.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_resent set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, no_of_gas_wells = :7, no_of_injection_wells = :8, no_of_oil_wells = :9, ppdm_guid = :10, remark = :11, resent_id = :12, short_name = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where pden_subtype = :21")
	if err != nil {
		return err
	}

	pden_resent.Row_changed_date = formatDate(pden_resent.Row_changed_date)
	pden_resent.Effective_date = formatDateString(pden_resent.Effective_date)
	pden_resent.Expiry_date = formatDateString(pden_resent.Expiry_date)
	pden_resent.Row_effective_date = formatDateString(pden_resent.Row_effective_date)
	pden_resent.Row_expiry_date = formatDateString(pden_resent.Row_expiry_date)






	rows, err := stmt.Exec(pden_resent.Pden_id, pden_resent.Pden_source, pden_resent.Active_ind, pden_resent.Effective_date, pden_resent.Expiry_date, pden_resent.Long_name, pden_resent.No_of_gas_wells, pden_resent.No_of_injection_wells, pden_resent.No_of_oil_wells, pden_resent.Ppdm_guid, pden_resent.Remark, pden_resent.Resent_id, pden_resent.Short_name, pden_resent.Row_changed_by, pden_resent.Row_changed_date, pden_resent.Row_created_by, pden_resent.Row_effective_date, pden_resent.Row_expiry_date, pden_resent.Row_quality, pden_resent.Pden_subtype)
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

func PatchPdenResent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_resent set "
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

func DeletePdenResent(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_resent dto.Pden_resent
	pden_resent.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_resent where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_resent.Pden_subtype)
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


