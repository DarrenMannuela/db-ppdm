package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellShow(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_show")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_show

	for rows.Next() {
		var well_show dto.Well_show
		if err := rows.Scan(&well_show.Uwi, &well_show.Source, &well_show.Show_type, &well_show.Show_obs_no, &well_show.Active_ind, &well_show.Base_depth, &well_show.Base_depth_ouom, &well_show.Base_strat_unit_id, &well_show.Effective_date, &well_show.Expiry_date, &well_show.Lithology_desc, &well_show.Ppdm_guid, &well_show.Remark, &well_show.Reservoir_ind, &well_show.Sample_type, &well_show.Show_symbol, &well_show.Source_document_id, &well_show.Strat_name_set_id, &well_show.Top_depth, &well_show.Top_depth_ouom, &well_show.Top_strat_unit_id, &well_show.Row_changed_by, &well_show.Row_changed_date, &well_show.Row_created_by, &well_show.Row_created_date, &well_show.Row_effective_date, &well_show.Row_expiry_date, &well_show.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_show to result
		result = append(result, well_show)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellShow(c *fiber.Ctx) error {
	var well_show dto.Well_show

	setDefaults(&well_show)

	if err := json.Unmarshal(c.Body(), &well_show); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_show values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	well_show.Row_created_date = formatDate(well_show.Row_created_date)
	well_show.Row_changed_date = formatDate(well_show.Row_changed_date)
	well_show.Effective_date = formatDateString(well_show.Effective_date)
	well_show.Expiry_date = formatDateString(well_show.Expiry_date)
	well_show.Row_effective_date = formatDateString(well_show.Row_effective_date)
	well_show.Row_expiry_date = formatDateString(well_show.Row_expiry_date)






	rows, err := stmt.Exec(well_show.Uwi, well_show.Source, well_show.Show_type, well_show.Show_obs_no, well_show.Active_ind, well_show.Base_depth, well_show.Base_depth_ouom, well_show.Base_strat_unit_id, well_show.Effective_date, well_show.Expiry_date, well_show.Lithology_desc, well_show.Ppdm_guid, well_show.Remark, well_show.Reservoir_ind, well_show.Sample_type, well_show.Show_symbol, well_show.Source_document_id, well_show.Strat_name_set_id, well_show.Top_depth, well_show.Top_depth_ouom, well_show.Top_strat_unit_id, well_show.Row_changed_by, well_show.Row_changed_date, well_show.Row_created_by, well_show.Row_created_date, well_show.Row_effective_date, well_show.Row_expiry_date, well_show.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellShow(c *fiber.Ctx) error {
	var well_show dto.Well_show

	setDefaults(&well_show)

	if err := json.Unmarshal(c.Body(), &well_show); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_show.Uwi = id

    if well_show.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_show set source = :1, show_type = :2, show_obs_no = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, base_strat_unit_id = :7, effective_date = :8, expiry_date = :9, lithology_desc = :10, ppdm_guid = :11, remark = :12, reservoir_ind = :13, sample_type = :14, show_symbol = :15, source_document_id = :16, strat_name_set_id = :17, top_depth = :18, top_depth_ouom = :19, top_strat_unit_id = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	well_show.Row_changed_date = formatDate(well_show.Row_changed_date)
	well_show.Effective_date = formatDateString(well_show.Effective_date)
	well_show.Expiry_date = formatDateString(well_show.Expiry_date)
	well_show.Row_effective_date = formatDateString(well_show.Row_effective_date)
	well_show.Row_expiry_date = formatDateString(well_show.Row_expiry_date)






	rows, err := stmt.Exec(well_show.Source, well_show.Show_type, well_show.Show_obs_no, well_show.Active_ind, well_show.Base_depth, well_show.Base_depth_ouom, well_show.Base_strat_unit_id, well_show.Effective_date, well_show.Expiry_date, well_show.Lithology_desc, well_show.Ppdm_guid, well_show.Remark, well_show.Reservoir_ind, well_show.Sample_type, well_show.Show_symbol, well_show.Source_document_id, well_show.Strat_name_set_id, well_show.Top_depth, well_show.Top_depth_ouom, well_show.Top_strat_unit_id, well_show.Row_changed_by, well_show.Row_changed_date, well_show.Row_created_by, well_show.Row_effective_date, well_show.Row_expiry_date, well_show.Row_quality, well_show.Uwi)
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

func PatchWellShow(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_show set "
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

func DeleteWellShow(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_show dto.Well_show
	well_show.Uwi = id

	stmt, err := db.Prepare("delete from well_show where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_show.Uwi)
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


