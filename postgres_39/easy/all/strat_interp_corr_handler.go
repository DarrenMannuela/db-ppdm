package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratInterpCorr(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_interp_corr")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_interp_corr

	for rows.Next() {
		var strat_interp_corr dto.Strat_interp_corr
		if err := rows.Scan(&strat_interp_corr.Correlation_id, &strat_interp_corr.Active_ind, &strat_interp_corr.Area_id, &strat_interp_corr.Area_type, &strat_interp_corr.Business_associate_id, &strat_interp_corr.Correlation_date, &strat_interp_corr.Effective_date, &strat_interp_corr.Expiry_date, &strat_interp_corr.Field_interp_id_1, &strat_interp_corr.Field_interp_id_2, &strat_interp_corr.Field_station_id_1, &strat_interp_corr.Field_station_id_2, &strat_interp_corr.Field_strat_name_set_1, &strat_interp_corr.Field_strat_name_set_2, &strat_interp_corr.Field_strat_unit_id_1, &strat_interp_corr.Field_strat_unit_id_2, &strat_interp_corr.Ppdm_guid, &strat_interp_corr.Project_id, &strat_interp_corr.Remark, &strat_interp_corr.Source, &strat_interp_corr.Source_document_id, &strat_interp_corr.Strat_correlation_criteria, &strat_interp_corr.Strat_correlation_quality, &strat_interp_corr.Strat_correlation_type, &strat_interp_corr.Strat_interpret_method, &strat_interp_corr.Uwi_1, &strat_interp_corr.Uwi_2, &strat_interp_corr.Uwi_interp_id_1, &strat_interp_corr.Uwi_interp_id_2, &strat_interp_corr.Uwi_strat_name_set_1, &strat_interp_corr.Uwi_strat_name_set_2, &strat_interp_corr.Uwi_strat_unit_id_1, &strat_interp_corr.Uwi_strat_unit_id_2, &strat_interp_corr.Row_changed_by, &strat_interp_corr.Row_changed_date, &strat_interp_corr.Row_created_by, &strat_interp_corr.Row_created_date, &strat_interp_corr.Row_effective_date, &strat_interp_corr.Row_expiry_date, &strat_interp_corr.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_interp_corr to result
		result = append(result, strat_interp_corr)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratInterpCorr(c *fiber.Ctx) error {
	var strat_interp_corr dto.Strat_interp_corr

	setDefaults(&strat_interp_corr)

	if err := json.Unmarshal(c.Body(), &strat_interp_corr); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_interp_corr values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	strat_interp_corr.Row_created_date = formatDate(strat_interp_corr.Row_created_date)
	strat_interp_corr.Row_changed_date = formatDate(strat_interp_corr.Row_changed_date)
	strat_interp_corr.Correlation_date = formatDateString(strat_interp_corr.Correlation_date)
	strat_interp_corr.Effective_date = formatDateString(strat_interp_corr.Effective_date)
	strat_interp_corr.Expiry_date = formatDateString(strat_interp_corr.Expiry_date)
	strat_interp_corr.Row_effective_date = formatDateString(strat_interp_corr.Row_effective_date)
	strat_interp_corr.Row_expiry_date = formatDateString(strat_interp_corr.Row_expiry_date)






	rows, err := stmt.Exec(strat_interp_corr.Correlation_id, strat_interp_corr.Active_ind, strat_interp_corr.Area_id, strat_interp_corr.Area_type, strat_interp_corr.Business_associate_id, strat_interp_corr.Correlation_date, strat_interp_corr.Effective_date, strat_interp_corr.Expiry_date, strat_interp_corr.Field_interp_id_1, strat_interp_corr.Field_interp_id_2, strat_interp_corr.Field_station_id_1, strat_interp_corr.Field_station_id_2, strat_interp_corr.Field_strat_name_set_1, strat_interp_corr.Field_strat_name_set_2, strat_interp_corr.Field_strat_unit_id_1, strat_interp_corr.Field_strat_unit_id_2, strat_interp_corr.Ppdm_guid, strat_interp_corr.Project_id, strat_interp_corr.Remark, strat_interp_corr.Source, strat_interp_corr.Source_document_id, strat_interp_corr.Strat_correlation_criteria, strat_interp_corr.Strat_correlation_quality, strat_interp_corr.Strat_correlation_type, strat_interp_corr.Strat_interpret_method, strat_interp_corr.Uwi_1, strat_interp_corr.Uwi_2, strat_interp_corr.Uwi_interp_id_1, strat_interp_corr.Uwi_interp_id_2, strat_interp_corr.Uwi_strat_name_set_1, strat_interp_corr.Uwi_strat_name_set_2, strat_interp_corr.Uwi_strat_unit_id_1, strat_interp_corr.Uwi_strat_unit_id_2, strat_interp_corr.Row_changed_by, strat_interp_corr.Row_changed_date, strat_interp_corr.Row_created_by, strat_interp_corr.Row_created_date, strat_interp_corr.Row_effective_date, strat_interp_corr.Row_expiry_date, strat_interp_corr.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratInterpCorr(c *fiber.Ctx) error {
	var strat_interp_corr dto.Strat_interp_corr

	setDefaults(&strat_interp_corr)

	if err := json.Unmarshal(c.Body(), &strat_interp_corr); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_interp_corr.Correlation_id = id

    if strat_interp_corr.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_interp_corr set active_ind = :1, area_id = :2, area_type = :3, business_associate_id = :4, correlation_date = :5, effective_date = :6, expiry_date = :7, field_interp_id_1 = :8, field_interp_id_2 = :9, field_station_id_1 = :10, field_station_id_2 = :11, field_strat_name_set_1 = :12, field_strat_name_set_2 = :13, field_strat_unit_id_1 = :14, field_strat_unit_id_2 = :15, ppdm_guid = :16, project_id = :17, remark = :18, source = :19, source_document_id = :20, strat_correlation_criteria = :21, strat_correlation_quality = :22, strat_correlation_type = :23, strat_interpret_method = :24, uwi_1 = :25, uwi_2 = :26, uwi_interp_id_1 = :27, uwi_interp_id_2 = :28, uwi_strat_name_set_1 = :29, uwi_strat_name_set_2 = :30, uwi_strat_unit_id_1 = :31, uwi_strat_unit_id_2 = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where correlation_id = :40")
	if err != nil {
		return err
	}

	strat_interp_corr.Row_changed_date = formatDate(strat_interp_corr.Row_changed_date)
	strat_interp_corr.Correlation_date = formatDateString(strat_interp_corr.Correlation_date)
	strat_interp_corr.Effective_date = formatDateString(strat_interp_corr.Effective_date)
	strat_interp_corr.Expiry_date = formatDateString(strat_interp_corr.Expiry_date)
	strat_interp_corr.Row_effective_date = formatDateString(strat_interp_corr.Row_effective_date)
	strat_interp_corr.Row_expiry_date = formatDateString(strat_interp_corr.Row_expiry_date)






	rows, err := stmt.Exec(strat_interp_corr.Active_ind, strat_interp_corr.Area_id, strat_interp_corr.Area_type, strat_interp_corr.Business_associate_id, strat_interp_corr.Correlation_date, strat_interp_corr.Effective_date, strat_interp_corr.Expiry_date, strat_interp_corr.Field_interp_id_1, strat_interp_corr.Field_interp_id_2, strat_interp_corr.Field_station_id_1, strat_interp_corr.Field_station_id_2, strat_interp_corr.Field_strat_name_set_1, strat_interp_corr.Field_strat_name_set_2, strat_interp_corr.Field_strat_unit_id_1, strat_interp_corr.Field_strat_unit_id_2, strat_interp_corr.Ppdm_guid, strat_interp_corr.Project_id, strat_interp_corr.Remark, strat_interp_corr.Source, strat_interp_corr.Source_document_id, strat_interp_corr.Strat_correlation_criteria, strat_interp_corr.Strat_correlation_quality, strat_interp_corr.Strat_correlation_type, strat_interp_corr.Strat_interpret_method, strat_interp_corr.Uwi_1, strat_interp_corr.Uwi_2, strat_interp_corr.Uwi_interp_id_1, strat_interp_corr.Uwi_interp_id_2, strat_interp_corr.Uwi_strat_name_set_1, strat_interp_corr.Uwi_strat_name_set_2, strat_interp_corr.Uwi_strat_unit_id_1, strat_interp_corr.Uwi_strat_unit_id_2, strat_interp_corr.Row_changed_by, strat_interp_corr.Row_changed_date, strat_interp_corr.Row_created_by, strat_interp_corr.Row_effective_date, strat_interp_corr.Row_expiry_date, strat_interp_corr.Row_quality, strat_interp_corr.Correlation_id)
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

func PatchStratInterpCorr(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_interp_corr set "
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
		} else if key == "correlation_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where correlation_id = :id"

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

func DeleteStratInterpCorr(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_interp_corr dto.Strat_interp_corr
	strat_interp_corr.Correlation_id = id

	stmt, err := db.Prepare("delete from strat_interp_corr where correlation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_interp_corr.Correlation_id)
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


