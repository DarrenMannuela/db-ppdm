package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCement(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_cement")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_cement

	for rows.Next() {
		var well_cement dto.Well_cement
		if err := rows.Scan(&well_cement.Uwi, &well_cement.Well_tube_source, &well_cement.Tubing_type, &well_cement.Tubing_obs_no, &well_cement.Cement_obs_no, &well_cement.Active_ind, &well_cement.Cement_amount, &well_cement.Cement_amount_ouom, &well_cement.Cement_amount_uom, &well_cement.Cement_ba_id, &well_cement.Cement_type, &well_cement.Effective_date, &well_cement.Expiry_date, &well_cement.Perforation_base_depth, &well_cement.Perforation_base_depth_ouom, &well_cement.Perforation_count, &well_cement.Perforation_per_uom, &well_cement.Perforation_top_depth, &well_cement.Perforation_top_depth_ouom, &well_cement.Ppdm_guid, &well_cement.Recement_ind, &well_cement.Remark, &well_cement.Source, &well_cement.Stage_no, &well_cement.Row_changed_by, &well_cement.Row_changed_date, &well_cement.Row_created_by, &well_cement.Row_created_date, &well_cement.Row_effective_date, &well_cement.Row_expiry_date, &well_cement.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_cement to result
		result = append(result, well_cement)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCement(c *fiber.Ctx) error {
	var well_cement dto.Well_cement

	setDefaults(&well_cement)

	if err := json.Unmarshal(c.Body(), &well_cement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_cement values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	well_cement.Row_created_date = formatDate(well_cement.Row_created_date)
	well_cement.Row_changed_date = formatDate(well_cement.Row_changed_date)
	well_cement.Effective_date = formatDateString(well_cement.Effective_date)
	well_cement.Expiry_date = formatDateString(well_cement.Expiry_date)
	well_cement.Row_effective_date = formatDateString(well_cement.Row_effective_date)
	well_cement.Row_expiry_date = formatDateString(well_cement.Row_expiry_date)






	rows, err := stmt.Exec(well_cement.Uwi, well_cement.Well_tube_source, well_cement.Tubing_type, well_cement.Tubing_obs_no, well_cement.Cement_obs_no, well_cement.Active_ind, well_cement.Cement_amount, well_cement.Cement_amount_ouom, well_cement.Cement_amount_uom, well_cement.Cement_ba_id, well_cement.Cement_type, well_cement.Effective_date, well_cement.Expiry_date, well_cement.Perforation_base_depth, well_cement.Perforation_base_depth_ouom, well_cement.Perforation_count, well_cement.Perforation_per_uom, well_cement.Perforation_top_depth, well_cement.Perforation_top_depth_ouom, well_cement.Ppdm_guid, well_cement.Recement_ind, well_cement.Remark, well_cement.Source, well_cement.Stage_no, well_cement.Row_changed_by, well_cement.Row_changed_date, well_cement.Row_created_by, well_cement.Row_created_date, well_cement.Row_effective_date, well_cement.Row_expiry_date, well_cement.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCement(c *fiber.Ctx) error {
	var well_cement dto.Well_cement

	setDefaults(&well_cement)

	if err := json.Unmarshal(c.Body(), &well_cement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_cement.Uwi = id

    if well_cement.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_cement set well_tube_source = :1, tubing_type = :2, tubing_obs_no = :3, cement_obs_no = :4, active_ind = :5, cement_amount = :6, cement_amount_ouom = :7, cement_amount_uom = :8, cement_ba_id = :9, cement_type = :10, effective_date = :11, expiry_date = :12, perforation_base_depth = :13, perforation_base_depth_ouom = :14, perforation_count = :15, perforation_per_uom = :16, perforation_top_depth = :17, perforation_top_depth_ouom = :18, ppdm_guid = :19, recement_ind = :20, remark = :21, source = :22, stage_no = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where uwi = :31")
	if err != nil {
		return err
	}

	well_cement.Row_changed_date = formatDate(well_cement.Row_changed_date)
	well_cement.Effective_date = formatDateString(well_cement.Effective_date)
	well_cement.Expiry_date = formatDateString(well_cement.Expiry_date)
	well_cement.Row_effective_date = formatDateString(well_cement.Row_effective_date)
	well_cement.Row_expiry_date = formatDateString(well_cement.Row_expiry_date)






	rows, err := stmt.Exec(well_cement.Well_tube_source, well_cement.Tubing_type, well_cement.Tubing_obs_no, well_cement.Cement_obs_no, well_cement.Active_ind, well_cement.Cement_amount, well_cement.Cement_amount_ouom, well_cement.Cement_amount_uom, well_cement.Cement_ba_id, well_cement.Cement_type, well_cement.Effective_date, well_cement.Expiry_date, well_cement.Perforation_base_depth, well_cement.Perforation_base_depth_ouom, well_cement.Perforation_count, well_cement.Perforation_per_uom, well_cement.Perforation_top_depth, well_cement.Perforation_top_depth_ouom, well_cement.Ppdm_guid, well_cement.Recement_ind, well_cement.Remark, well_cement.Source, well_cement.Stage_no, well_cement.Row_changed_by, well_cement.Row_changed_date, well_cement.Row_created_by, well_cement.Row_effective_date, well_cement.Row_expiry_date, well_cement.Row_quality, well_cement.Uwi)
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

func PatchWellCement(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_cement set "
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

func DeleteWellCement(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_cement dto.Well_cement
	well_cement.Uwi = id

	stmt, err := db.Prepare("delete from well_cement where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_cement.Uwi)
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


