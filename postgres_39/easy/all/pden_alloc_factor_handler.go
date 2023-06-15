package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenAllocFactor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_alloc_factor")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_alloc_factor

	for rows.Next() {
		var pden_alloc_factor dto.Pden_alloc_factor
		if err := rows.Scan(&pden_alloc_factor.From_pden_subtype, &pden_alloc_factor.From_pden_id, &pden_alloc_factor.From_pden_source, &pden_alloc_factor.To_pden_subtype, &pden_alloc_factor.To_pden_id, &pden_alloc_factor.To_pden_source, &pden_alloc_factor.Allocation_obs_no, &pden_alloc_factor.Active_ind, &pden_alloc_factor.Activity_type, &pden_alloc_factor.Allocation_factor, &pden_alloc_factor.Allocation_type, &pden_alloc_factor.Business_rule, &pden_alloc_factor.Calculated_from, &pden_alloc_factor.Effective_date, &pden_alloc_factor.Expiry_date, &pden_alloc_factor.Period_type, &pden_alloc_factor.Ppdm_guid, &pden_alloc_factor.Product_type, &pden_alloc_factor.Remark, &pden_alloc_factor.Source, &pden_alloc_factor.Volume_method, &pden_alloc_factor.Row_changed_by, &pden_alloc_factor.Row_changed_date, &pden_alloc_factor.Row_created_by, &pden_alloc_factor.Row_created_date, &pden_alloc_factor.Row_effective_date, &pden_alloc_factor.Row_expiry_date, &pden_alloc_factor.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_alloc_factor to result
		result = append(result, pden_alloc_factor)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenAllocFactor(c *fiber.Ctx) error {
	var pden_alloc_factor dto.Pden_alloc_factor

	setDefaults(&pden_alloc_factor)

	if err := json.Unmarshal(c.Body(), &pden_alloc_factor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_alloc_factor values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	pden_alloc_factor.Row_created_date = formatDate(pden_alloc_factor.Row_created_date)
	pden_alloc_factor.Row_changed_date = formatDate(pden_alloc_factor.Row_changed_date)
	pden_alloc_factor.Effective_date = formatDateString(pden_alloc_factor.Effective_date)
	pden_alloc_factor.Expiry_date = formatDateString(pden_alloc_factor.Expiry_date)
	pden_alloc_factor.Row_effective_date = formatDateString(pden_alloc_factor.Row_effective_date)
	pden_alloc_factor.Row_expiry_date = formatDateString(pden_alloc_factor.Row_expiry_date)






	rows, err := stmt.Exec(pden_alloc_factor.From_pden_subtype, pden_alloc_factor.From_pden_id, pden_alloc_factor.From_pden_source, pden_alloc_factor.To_pden_subtype, pden_alloc_factor.To_pden_id, pden_alloc_factor.To_pden_source, pden_alloc_factor.Allocation_obs_no, pden_alloc_factor.Active_ind, pden_alloc_factor.Activity_type, pden_alloc_factor.Allocation_factor, pden_alloc_factor.Allocation_type, pden_alloc_factor.Business_rule, pden_alloc_factor.Calculated_from, pden_alloc_factor.Effective_date, pden_alloc_factor.Expiry_date, pden_alloc_factor.Period_type, pden_alloc_factor.Ppdm_guid, pden_alloc_factor.Product_type, pden_alloc_factor.Remark, pden_alloc_factor.Source, pden_alloc_factor.Volume_method, pden_alloc_factor.Row_changed_by, pden_alloc_factor.Row_changed_date, pden_alloc_factor.Row_created_by, pden_alloc_factor.Row_created_date, pden_alloc_factor.Row_effective_date, pden_alloc_factor.Row_expiry_date, pden_alloc_factor.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenAllocFactor(c *fiber.Ctx) error {
	var pden_alloc_factor dto.Pden_alloc_factor

	setDefaults(&pden_alloc_factor)

	if err := json.Unmarshal(c.Body(), &pden_alloc_factor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_alloc_factor.From_pden_subtype = id

    if pden_alloc_factor.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_alloc_factor set from_pden_id = :1, from_pden_source = :2, to_pden_subtype = :3, to_pden_id = :4, to_pden_source = :5, allocation_obs_no = :6, active_ind = :7, activity_type = :8, allocation_factor = :9, allocation_type = :10, business_rule = :11, calculated_from = :12, effective_date = :13, expiry_date = :14, period_type = :15, ppdm_guid = :16, product_type = :17, remark = :18, source = :19, volume_method = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where from_pden_subtype = :28")
	if err != nil {
		return err
	}

	pden_alloc_factor.Row_changed_date = formatDate(pden_alloc_factor.Row_changed_date)
	pden_alloc_factor.Effective_date = formatDateString(pden_alloc_factor.Effective_date)
	pden_alloc_factor.Expiry_date = formatDateString(pden_alloc_factor.Expiry_date)
	pden_alloc_factor.Row_effective_date = formatDateString(pden_alloc_factor.Row_effective_date)
	pden_alloc_factor.Row_expiry_date = formatDateString(pden_alloc_factor.Row_expiry_date)






	rows, err := stmt.Exec(pden_alloc_factor.From_pden_id, pden_alloc_factor.From_pden_source, pden_alloc_factor.To_pden_subtype, pden_alloc_factor.To_pden_id, pden_alloc_factor.To_pden_source, pden_alloc_factor.Allocation_obs_no, pden_alloc_factor.Active_ind, pden_alloc_factor.Activity_type, pden_alloc_factor.Allocation_factor, pden_alloc_factor.Allocation_type, pden_alloc_factor.Business_rule, pden_alloc_factor.Calculated_from, pden_alloc_factor.Effective_date, pden_alloc_factor.Expiry_date, pden_alloc_factor.Period_type, pden_alloc_factor.Ppdm_guid, pden_alloc_factor.Product_type, pden_alloc_factor.Remark, pden_alloc_factor.Source, pden_alloc_factor.Volume_method, pden_alloc_factor.Row_changed_by, pden_alloc_factor.Row_changed_date, pden_alloc_factor.Row_created_by, pden_alloc_factor.Row_effective_date, pden_alloc_factor.Row_expiry_date, pden_alloc_factor.Row_quality, pden_alloc_factor.From_pden_subtype)
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

func PatchPdenAllocFactor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_alloc_factor set "
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
	query += " where from_pden_subtype = :id"

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

func DeletePdenAllocFactor(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_alloc_factor dto.Pden_alloc_factor
	pden_alloc_factor.From_pden_subtype = id

	stmt, err := db.Prepare("delete from pden_alloc_factor where from_pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_alloc_factor.From_pden_subtype)
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


