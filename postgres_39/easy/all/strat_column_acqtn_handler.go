package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratColumnAcqtn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_column_acqtn")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_column_acqtn

	for rows.Next() {
		var strat_column_acqtn dto.Strat_column_acqtn
		if err := rows.Scan(&strat_column_acqtn.Strat_column_id, &strat_column_acqtn.Strat_column_source, &strat_column_acqtn.Strat_name_set_id, &strat_column_acqtn.Strat_unit_id, &strat_column_acqtn.Interp_id, &strat_column_acqtn.Strat_column_acqtn_id, &strat_column_acqtn.Active_ind, &strat_column_acqtn.Core_id, &strat_column_acqtn.Core_source, &strat_column_acqtn.Description, &strat_column_acqtn.Effective_date, &strat_column_acqtn.Expiry_date, &strat_column_acqtn.Field_interp_id, &strat_column_acqtn.Field_station_id, &strat_column_acqtn.Field_strat_unit_id, &strat_column_acqtn.Information_item_id, &strat_column_acqtn.Info_item_subtype, &strat_column_acqtn.Log_curve_id, &strat_column_acqtn.Ppdm_guid, &strat_column_acqtn.Project_id, &strat_column_acqtn.Remark, &strat_column_acqtn.Seis_set_id, &strat_column_acqtn.Seis_set_subtype, &strat_column_acqtn.Source, &strat_column_acqtn.Uwi, &strat_column_acqtn.Ws_interp_id, &strat_column_acqtn.Ws_strat_unit_id, &strat_column_acqtn.Row_changed_by, &strat_column_acqtn.Row_changed_date, &strat_column_acqtn.Row_created_by, &strat_column_acqtn.Row_created_date, &strat_column_acqtn.Row_effective_date, &strat_column_acqtn.Row_expiry_date, &strat_column_acqtn.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_column_acqtn to result
		result = append(result, strat_column_acqtn)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratColumnAcqtn(c *fiber.Ctx) error {
	var strat_column_acqtn dto.Strat_column_acqtn

	setDefaults(&strat_column_acqtn)

	if err := json.Unmarshal(c.Body(), &strat_column_acqtn); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_column_acqtn values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	strat_column_acqtn.Row_created_date = formatDate(strat_column_acqtn.Row_created_date)
	strat_column_acqtn.Row_changed_date = formatDate(strat_column_acqtn.Row_changed_date)
	strat_column_acqtn.Effective_date = formatDateString(strat_column_acqtn.Effective_date)
	strat_column_acqtn.Expiry_date = formatDateString(strat_column_acqtn.Expiry_date)
	strat_column_acqtn.Row_effective_date = formatDateString(strat_column_acqtn.Row_effective_date)
	strat_column_acqtn.Row_expiry_date = formatDateString(strat_column_acqtn.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_acqtn.Strat_column_id, strat_column_acqtn.Strat_column_source, strat_column_acqtn.Strat_name_set_id, strat_column_acqtn.Strat_unit_id, strat_column_acqtn.Interp_id, strat_column_acqtn.Strat_column_acqtn_id, strat_column_acqtn.Active_ind, strat_column_acqtn.Core_id, strat_column_acqtn.Core_source, strat_column_acqtn.Description, strat_column_acqtn.Effective_date, strat_column_acqtn.Expiry_date, strat_column_acqtn.Field_interp_id, strat_column_acqtn.Field_station_id, strat_column_acqtn.Field_strat_unit_id, strat_column_acqtn.Information_item_id, strat_column_acqtn.Info_item_subtype, strat_column_acqtn.Log_curve_id, strat_column_acqtn.Ppdm_guid, strat_column_acqtn.Project_id, strat_column_acqtn.Remark, strat_column_acqtn.Seis_set_id, strat_column_acqtn.Seis_set_subtype, strat_column_acqtn.Source, strat_column_acqtn.Uwi, strat_column_acqtn.Ws_interp_id, strat_column_acqtn.Ws_strat_unit_id, strat_column_acqtn.Row_changed_by, strat_column_acqtn.Row_changed_date, strat_column_acqtn.Row_created_by, strat_column_acqtn.Row_created_date, strat_column_acqtn.Row_effective_date, strat_column_acqtn.Row_expiry_date, strat_column_acqtn.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratColumnAcqtn(c *fiber.Ctx) error {
	var strat_column_acqtn dto.Strat_column_acqtn

	setDefaults(&strat_column_acqtn)

	if err := json.Unmarshal(c.Body(), &strat_column_acqtn); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_column_acqtn.Strat_column_id = id

    if strat_column_acqtn.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_column_acqtn set strat_column_source = :1, strat_name_set_id = :2, strat_unit_id = :3, interp_id = :4, strat_column_acqtn_id = :5, active_ind = :6, core_id = :7, core_source = :8, description = :9, effective_date = :10, expiry_date = :11, field_interp_id = :12, field_station_id = :13, field_strat_unit_id = :14, information_item_id = :15, info_item_subtype = :16, log_curve_id = :17, ppdm_guid = :18, project_id = :19, remark = :20, seis_set_id = :21, seis_set_subtype = :22, source = :23, uwi = :24, ws_interp_id = :25, ws_strat_unit_id = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where strat_column_id = :34")
	if err != nil {
		return err
	}

	strat_column_acqtn.Row_changed_date = formatDate(strat_column_acqtn.Row_changed_date)
	strat_column_acqtn.Effective_date = formatDateString(strat_column_acqtn.Effective_date)
	strat_column_acqtn.Expiry_date = formatDateString(strat_column_acqtn.Expiry_date)
	strat_column_acqtn.Row_effective_date = formatDateString(strat_column_acqtn.Row_effective_date)
	strat_column_acqtn.Row_expiry_date = formatDateString(strat_column_acqtn.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_acqtn.Strat_column_source, strat_column_acqtn.Strat_name_set_id, strat_column_acqtn.Strat_unit_id, strat_column_acqtn.Interp_id, strat_column_acqtn.Strat_column_acqtn_id, strat_column_acqtn.Active_ind, strat_column_acqtn.Core_id, strat_column_acqtn.Core_source, strat_column_acqtn.Description, strat_column_acqtn.Effective_date, strat_column_acqtn.Expiry_date, strat_column_acqtn.Field_interp_id, strat_column_acqtn.Field_station_id, strat_column_acqtn.Field_strat_unit_id, strat_column_acqtn.Information_item_id, strat_column_acqtn.Info_item_subtype, strat_column_acqtn.Log_curve_id, strat_column_acqtn.Ppdm_guid, strat_column_acqtn.Project_id, strat_column_acqtn.Remark, strat_column_acqtn.Seis_set_id, strat_column_acqtn.Seis_set_subtype, strat_column_acqtn.Source, strat_column_acqtn.Uwi, strat_column_acqtn.Ws_interp_id, strat_column_acqtn.Ws_strat_unit_id, strat_column_acqtn.Row_changed_by, strat_column_acqtn.Row_changed_date, strat_column_acqtn.Row_created_by, strat_column_acqtn.Row_effective_date, strat_column_acqtn.Row_expiry_date, strat_column_acqtn.Row_quality, strat_column_acqtn.Strat_column_id)
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

func PatchStratColumnAcqtn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_column_acqtn set "
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
	query += " where strat_column_id = :id"

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

func DeleteStratColumnAcqtn(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_column_acqtn dto.Strat_column_acqtn
	strat_column_acqtn.Strat_column_id = id

	stmt, err := db.Prepare("delete from strat_column_acqtn where strat_column_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_column_acqtn.Strat_column_id)
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


