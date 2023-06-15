package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdStringFormation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_string_formation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_string_formation

	for rows.Next() {
		var prod_string_formation dto.Prod_string_formation
		if err := rows.Scan(&prod_string_formation.Uwi, &prod_string_formation.Prod_string_source, &prod_string_formation.String_id, &prod_string_formation.Pr_str_form_obs_no, &prod_string_formation.Active_ind, &prod_string_formation.Allocation_factor, &prod_string_formation.Allocation_type, &prod_string_formation.Base_depth, &prod_string_formation.Base_depth_ouom, &prod_string_formation.Completion_obs_no, &prod_string_formation.Current_status, &prod_string_formation.Current_status_date, &prod_string_formation.Effective_date, &prod_string_formation.Expiry_date, &prod_string_formation.Pool_id, &prod_string_formation.Ppdm_guid, &prod_string_formation.Remark, &prod_string_formation.Source, &prod_string_formation.Status_type, &prod_string_formation.Strat_name_set_id, &prod_string_formation.Strat_unit_id, &prod_string_formation.Top_depth, &prod_string_formation.Top_depth_ouom, &prod_string_formation.Row_changed_by, &prod_string_formation.Row_changed_date, &prod_string_formation.Row_created_by, &prod_string_formation.Row_created_date, &prod_string_formation.Row_effective_date, &prod_string_formation.Row_expiry_date, &prod_string_formation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_string_formation to result
		result = append(result, prod_string_formation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdStringFormation(c *fiber.Ctx) error {
	var prod_string_formation dto.Prod_string_formation

	setDefaults(&prod_string_formation)

	if err := json.Unmarshal(c.Body(), &prod_string_formation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_string_formation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	prod_string_formation.Row_created_date = formatDate(prod_string_formation.Row_created_date)
	prod_string_formation.Row_changed_date = formatDate(prod_string_formation.Row_changed_date)
	prod_string_formation.Current_status_date = formatDateString(prod_string_formation.Current_status_date)
	prod_string_formation.Effective_date = formatDateString(prod_string_formation.Effective_date)
	prod_string_formation.Expiry_date = formatDateString(prod_string_formation.Expiry_date)
	prod_string_formation.Row_effective_date = formatDateString(prod_string_formation.Row_effective_date)
	prod_string_formation.Row_expiry_date = formatDateString(prod_string_formation.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_formation.Uwi, prod_string_formation.Prod_string_source, prod_string_formation.String_id, prod_string_formation.Pr_str_form_obs_no, prod_string_formation.Active_ind, prod_string_formation.Allocation_factor, prod_string_formation.Allocation_type, prod_string_formation.Base_depth, prod_string_formation.Base_depth_ouom, prod_string_formation.Completion_obs_no, prod_string_formation.Current_status, prod_string_formation.Current_status_date, prod_string_formation.Effective_date, prod_string_formation.Expiry_date, prod_string_formation.Pool_id, prod_string_formation.Ppdm_guid, prod_string_formation.Remark, prod_string_formation.Source, prod_string_formation.Status_type, prod_string_formation.Strat_name_set_id, prod_string_formation.Strat_unit_id, prod_string_formation.Top_depth, prod_string_formation.Top_depth_ouom, prod_string_formation.Row_changed_by, prod_string_formation.Row_changed_date, prod_string_formation.Row_created_by, prod_string_formation.Row_created_date, prod_string_formation.Row_effective_date, prod_string_formation.Row_expiry_date, prod_string_formation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdStringFormation(c *fiber.Ctx) error {
	var prod_string_formation dto.Prod_string_formation

	setDefaults(&prod_string_formation)

	if err := json.Unmarshal(c.Body(), &prod_string_formation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_string_formation.Uwi = id

    if prod_string_formation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_string_formation set prod_string_source = :1, string_id = :2, pr_str_form_obs_no = :3, active_ind = :4, allocation_factor = :5, allocation_type = :6, base_depth = :7, base_depth_ouom = :8, completion_obs_no = :9, current_status = :10, current_status_date = :11, effective_date = :12, expiry_date = :13, pool_id = :14, ppdm_guid = :15, remark = :16, source = :17, status_type = :18, strat_name_set_id = :19, strat_unit_id = :20, top_depth = :21, top_depth_ouom = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where uwi = :30")
	if err != nil {
		return err
	}

	prod_string_formation.Row_changed_date = formatDate(prod_string_formation.Row_changed_date)
	prod_string_formation.Current_status_date = formatDateString(prod_string_formation.Current_status_date)
	prod_string_formation.Effective_date = formatDateString(prod_string_formation.Effective_date)
	prod_string_formation.Expiry_date = formatDateString(prod_string_formation.Expiry_date)
	prod_string_formation.Row_effective_date = formatDateString(prod_string_formation.Row_effective_date)
	prod_string_formation.Row_expiry_date = formatDateString(prod_string_formation.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_formation.Prod_string_source, prod_string_formation.String_id, prod_string_formation.Pr_str_form_obs_no, prod_string_formation.Active_ind, prod_string_formation.Allocation_factor, prod_string_formation.Allocation_type, prod_string_formation.Base_depth, prod_string_formation.Base_depth_ouom, prod_string_formation.Completion_obs_no, prod_string_formation.Current_status, prod_string_formation.Current_status_date, prod_string_formation.Effective_date, prod_string_formation.Expiry_date, prod_string_formation.Pool_id, prod_string_formation.Ppdm_guid, prod_string_formation.Remark, prod_string_formation.Source, prod_string_formation.Status_type, prod_string_formation.Strat_name_set_id, prod_string_formation.Strat_unit_id, prod_string_formation.Top_depth, prod_string_formation.Top_depth_ouom, prod_string_formation.Row_changed_by, prod_string_formation.Row_changed_date, prod_string_formation.Row_created_by, prod_string_formation.Row_effective_date, prod_string_formation.Row_expiry_date, prod_string_formation.Row_quality, prod_string_formation.Uwi)
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

func PatchProdStringFormation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_string_formation set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteProdStringFormation(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_string_formation dto.Prod_string_formation
	prod_string_formation.Uwi = id

	stmt, err := db.Prepare("delete from prod_string_formation where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_string_formation.Uwi)
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


