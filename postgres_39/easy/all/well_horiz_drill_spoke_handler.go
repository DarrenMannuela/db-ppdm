package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellHorizDrillSpoke(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_horiz_drill_spoke")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_horiz_drill_spoke

	for rows.Next() {
		var well_horiz_drill_spoke dto.Well_horiz_drill_spoke
		if err := rows.Scan(&well_horiz_drill_spoke.Uwi, &well_horiz_drill_spoke.Source, &well_horiz_drill_spoke.Kickoff_point_obs_no, &well_horiz_drill_spoke.Spoke_obs_no, &well_horiz_drill_spoke.Active_ind, &well_horiz_drill_spoke.Effective_date, &well_horiz_drill_spoke.Expiry_date, &well_horiz_drill_spoke.Lateral_hole_id, &well_horiz_drill_spoke.Node_id, &well_horiz_drill_spoke.Ppdm_guid, &well_horiz_drill_spoke.Remark, &well_horiz_drill_spoke.Spoke_length, &well_horiz_drill_spoke.Spoke_length_ouom, &well_horiz_drill_spoke.Spoke_md, &well_horiz_drill_spoke.Spoke_md_ouom, &well_horiz_drill_spoke.Spoke_tvd, &well_horiz_drill_spoke.Spoke_tvd_ouom, &well_horiz_drill_spoke.Row_changed_by, &well_horiz_drill_spoke.Row_changed_date, &well_horiz_drill_spoke.Row_created_by, &well_horiz_drill_spoke.Row_created_date, &well_horiz_drill_spoke.Row_effective_date, &well_horiz_drill_spoke.Row_expiry_date, &well_horiz_drill_spoke.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_horiz_drill_spoke to result
		result = append(result, well_horiz_drill_spoke)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellHorizDrillSpoke(c *fiber.Ctx) error {
	var well_horiz_drill_spoke dto.Well_horiz_drill_spoke

	setDefaults(&well_horiz_drill_spoke)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_spoke); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_horiz_drill_spoke values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	well_horiz_drill_spoke.Row_created_date = formatDate(well_horiz_drill_spoke.Row_created_date)
	well_horiz_drill_spoke.Row_changed_date = formatDate(well_horiz_drill_spoke.Row_changed_date)
	well_horiz_drill_spoke.Effective_date = formatDateString(well_horiz_drill_spoke.Effective_date)
	well_horiz_drill_spoke.Expiry_date = formatDateString(well_horiz_drill_spoke.Expiry_date)
	well_horiz_drill_spoke.Row_effective_date = formatDateString(well_horiz_drill_spoke.Row_effective_date)
	well_horiz_drill_spoke.Row_expiry_date = formatDateString(well_horiz_drill_spoke.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_spoke.Uwi, well_horiz_drill_spoke.Source, well_horiz_drill_spoke.Kickoff_point_obs_no, well_horiz_drill_spoke.Spoke_obs_no, well_horiz_drill_spoke.Active_ind, well_horiz_drill_spoke.Effective_date, well_horiz_drill_spoke.Expiry_date, well_horiz_drill_spoke.Lateral_hole_id, well_horiz_drill_spoke.Node_id, well_horiz_drill_spoke.Ppdm_guid, well_horiz_drill_spoke.Remark, well_horiz_drill_spoke.Spoke_length, well_horiz_drill_spoke.Spoke_length_ouom, well_horiz_drill_spoke.Spoke_md, well_horiz_drill_spoke.Spoke_md_ouom, well_horiz_drill_spoke.Spoke_tvd, well_horiz_drill_spoke.Spoke_tvd_ouom, well_horiz_drill_spoke.Row_changed_by, well_horiz_drill_spoke.Row_changed_date, well_horiz_drill_spoke.Row_created_by, well_horiz_drill_spoke.Row_created_date, well_horiz_drill_spoke.Row_effective_date, well_horiz_drill_spoke.Row_expiry_date, well_horiz_drill_spoke.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellHorizDrillSpoke(c *fiber.Ctx) error {
	var well_horiz_drill_spoke dto.Well_horiz_drill_spoke

	setDefaults(&well_horiz_drill_spoke)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_spoke); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_horiz_drill_spoke.Uwi = id

    if well_horiz_drill_spoke.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_horiz_drill_spoke set source = :1, kickoff_point_obs_no = :2, spoke_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, lateral_hole_id = :7, node_id = :8, ppdm_guid = :9, remark = :10, spoke_length = :11, spoke_length_ouom = :12, spoke_md = :13, spoke_md_ouom = :14, spoke_tvd = :15, spoke_tvd_ouom = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where uwi = :24")
	if err != nil {
		return err
	}

	well_horiz_drill_spoke.Row_changed_date = formatDate(well_horiz_drill_spoke.Row_changed_date)
	well_horiz_drill_spoke.Effective_date = formatDateString(well_horiz_drill_spoke.Effective_date)
	well_horiz_drill_spoke.Expiry_date = formatDateString(well_horiz_drill_spoke.Expiry_date)
	well_horiz_drill_spoke.Row_effective_date = formatDateString(well_horiz_drill_spoke.Row_effective_date)
	well_horiz_drill_spoke.Row_expiry_date = formatDateString(well_horiz_drill_spoke.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_spoke.Source, well_horiz_drill_spoke.Kickoff_point_obs_no, well_horiz_drill_spoke.Spoke_obs_no, well_horiz_drill_spoke.Active_ind, well_horiz_drill_spoke.Effective_date, well_horiz_drill_spoke.Expiry_date, well_horiz_drill_spoke.Lateral_hole_id, well_horiz_drill_spoke.Node_id, well_horiz_drill_spoke.Ppdm_guid, well_horiz_drill_spoke.Remark, well_horiz_drill_spoke.Spoke_length, well_horiz_drill_spoke.Spoke_length_ouom, well_horiz_drill_spoke.Spoke_md, well_horiz_drill_spoke.Spoke_md_ouom, well_horiz_drill_spoke.Spoke_tvd, well_horiz_drill_spoke.Spoke_tvd_ouom, well_horiz_drill_spoke.Row_changed_by, well_horiz_drill_spoke.Row_changed_date, well_horiz_drill_spoke.Row_created_by, well_horiz_drill_spoke.Row_effective_date, well_horiz_drill_spoke.Row_expiry_date, well_horiz_drill_spoke.Row_quality, well_horiz_drill_spoke.Uwi)
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

func PatchWellHorizDrillSpoke(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_horiz_drill_spoke set "
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

func DeleteWellHorizDrillSpoke(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_horiz_drill_spoke dto.Well_horiz_drill_spoke
	well_horiz_drill_spoke.Uwi = id

	stmt, err := db.Prepare("delete from well_horiz_drill_spoke where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_horiz_drill_spoke.Uwi)
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


