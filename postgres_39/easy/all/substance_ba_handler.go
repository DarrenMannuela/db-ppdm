package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_ba

	for rows.Next() {
		var substance_ba dto.Substance_ba
		if err := rows.Scan(&substance_ba.Substance_id, &substance_ba.Anl_source, &substance_ba.Ba_obs_no, &substance_ba.Active_ind, &substance_ba.Ba_role_type, &substance_ba.Cas_number, &substance_ba.Effective_date, &substance_ba.Expiry_date, &substance_ba.Location, &substance_ba.Manufacturer, &substance_ba.Ppdm_guid, &substance_ba.Price, &substance_ba.Price_ouom, &substance_ba.Problem_ind, &substance_ba.Purchase_quantity, &substance_ba.Purchase_quantity_ouom, &substance_ba.Purchase_quantity_type, &substance_ba.Remark, &substance_ba.Source, &substance_ba.Step_seq_no, &substance_ba.Substance_ba_id, &substance_ba.Supplier, &substance_ba.Row_changed_by, &substance_ba.Row_changed_date, &substance_ba.Row_created_by, &substance_ba.Row_created_date, &substance_ba.Row_effective_date, &substance_ba.Row_expiry_date, &substance_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_ba to result
		result = append(result, substance_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceBa(c *fiber.Ctx) error {
	var substance_ba dto.Substance_ba

	setDefaults(&substance_ba)

	if err := json.Unmarshal(c.Body(), &substance_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	substance_ba.Row_created_date = formatDate(substance_ba.Row_created_date)
	substance_ba.Row_changed_date = formatDate(substance_ba.Row_changed_date)
	substance_ba.Effective_date = formatDateString(substance_ba.Effective_date)
	substance_ba.Expiry_date = formatDateString(substance_ba.Expiry_date)
	substance_ba.Row_effective_date = formatDateString(substance_ba.Row_effective_date)
	substance_ba.Row_expiry_date = formatDateString(substance_ba.Row_expiry_date)






	rows, err := stmt.Exec(substance_ba.Substance_id, substance_ba.Anl_source, substance_ba.Ba_obs_no, substance_ba.Active_ind, substance_ba.Ba_role_type, substance_ba.Cas_number, substance_ba.Effective_date, substance_ba.Expiry_date, substance_ba.Location, substance_ba.Manufacturer, substance_ba.Ppdm_guid, substance_ba.Price, substance_ba.Price_ouom, substance_ba.Problem_ind, substance_ba.Purchase_quantity, substance_ba.Purchase_quantity_ouom, substance_ba.Purchase_quantity_type, substance_ba.Remark, substance_ba.Source, substance_ba.Step_seq_no, substance_ba.Substance_ba_id, substance_ba.Supplier, substance_ba.Row_changed_by, substance_ba.Row_changed_date, substance_ba.Row_created_by, substance_ba.Row_created_date, substance_ba.Row_effective_date, substance_ba.Row_expiry_date, substance_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceBa(c *fiber.Ctx) error {
	var substance_ba dto.Substance_ba

	setDefaults(&substance_ba)

	if err := json.Unmarshal(c.Body(), &substance_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_ba.Substance_id = id

    if substance_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_ba set anl_source = :1, ba_obs_no = :2, active_ind = :3, ba_role_type = :4, cas_number = :5, effective_date = :6, expiry_date = :7, location = :8, manufacturer = :9, ppdm_guid = :10, price = :11, price_ouom = :12, problem_ind = :13, purchase_quantity = :14, purchase_quantity_ouom = :15, purchase_quantity_type = :16, remark = :17, source = :18, step_seq_no = :19, substance_ba_id = :20, supplier = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where substance_id = :29")
	if err != nil {
		return err
	}

	substance_ba.Row_changed_date = formatDate(substance_ba.Row_changed_date)
	substance_ba.Effective_date = formatDateString(substance_ba.Effective_date)
	substance_ba.Expiry_date = formatDateString(substance_ba.Expiry_date)
	substance_ba.Row_effective_date = formatDateString(substance_ba.Row_effective_date)
	substance_ba.Row_expiry_date = formatDateString(substance_ba.Row_expiry_date)






	rows, err := stmt.Exec(substance_ba.Anl_source, substance_ba.Ba_obs_no, substance_ba.Active_ind, substance_ba.Ba_role_type, substance_ba.Cas_number, substance_ba.Effective_date, substance_ba.Expiry_date, substance_ba.Location, substance_ba.Manufacturer, substance_ba.Ppdm_guid, substance_ba.Price, substance_ba.Price_ouom, substance_ba.Problem_ind, substance_ba.Purchase_quantity, substance_ba.Purchase_quantity_ouom, substance_ba.Purchase_quantity_type, substance_ba.Remark, substance_ba.Source, substance_ba.Step_seq_no, substance_ba.Substance_ba_id, substance_ba.Supplier, substance_ba.Row_changed_by, substance_ba.Row_changed_date, substance_ba.Row_created_by, substance_ba.Row_effective_date, substance_ba.Row_expiry_date, substance_ba.Row_quality, substance_ba.Substance_id)
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

func PatchSubstanceBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_ba set "
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
	query += " where substance_id = :id"

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

func DeleteSubstanceBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_ba dto.Substance_ba
	substance_ba.Substance_id = id

	stmt, err := db.Prepare("delete from substance_ba where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_ba.Substance_id)
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


