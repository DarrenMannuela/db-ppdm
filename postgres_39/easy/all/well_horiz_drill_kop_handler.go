package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellHorizDrillKop(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_horiz_drill_kop")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_horiz_drill_kop

	for rows.Next() {
		var well_horiz_drill_kop dto.Well_horiz_drill_kop
		if err := rows.Scan(&well_horiz_drill_kop.Uwi, &well_horiz_drill_kop.Source, &well_horiz_drill_kop.Kickoff_point_obs_no, &well_horiz_drill_kop.Active_ind, &well_horiz_drill_kop.Effective_date, &well_horiz_drill_kop.Expiry_date, &well_horiz_drill_kop.Kickoff_point_md, &well_horiz_drill_kop.Kickoff_point_md_ouom, &well_horiz_drill_kop.Kickoff_point_tvd, &well_horiz_drill_kop.Kickoff_point_tvd_ouom, &well_horiz_drill_kop.Lateral_hole_id, &well_horiz_drill_kop.Node_id, &well_horiz_drill_kop.Ppdm_guid, &well_horiz_drill_kop.Remark, &well_horiz_drill_kop.Row_changed_by, &well_horiz_drill_kop.Row_changed_date, &well_horiz_drill_kop.Row_created_by, &well_horiz_drill_kop.Row_created_date, &well_horiz_drill_kop.Row_effective_date, &well_horiz_drill_kop.Row_expiry_date, &well_horiz_drill_kop.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_horiz_drill_kop to result
		result = append(result, well_horiz_drill_kop)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellHorizDrillKop(c *fiber.Ctx) error {
	var well_horiz_drill_kop dto.Well_horiz_drill_kop

	setDefaults(&well_horiz_drill_kop)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_kop); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_horiz_drill_kop values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_horiz_drill_kop.Row_created_date = formatDate(well_horiz_drill_kop.Row_created_date)
	well_horiz_drill_kop.Row_changed_date = formatDate(well_horiz_drill_kop.Row_changed_date)
	well_horiz_drill_kop.Effective_date = formatDateString(well_horiz_drill_kop.Effective_date)
	well_horiz_drill_kop.Expiry_date = formatDateString(well_horiz_drill_kop.Expiry_date)
	well_horiz_drill_kop.Row_effective_date = formatDateString(well_horiz_drill_kop.Row_effective_date)
	well_horiz_drill_kop.Row_expiry_date = formatDateString(well_horiz_drill_kop.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_kop.Uwi, well_horiz_drill_kop.Source, well_horiz_drill_kop.Kickoff_point_obs_no, well_horiz_drill_kop.Active_ind, well_horiz_drill_kop.Effective_date, well_horiz_drill_kop.Expiry_date, well_horiz_drill_kop.Kickoff_point_md, well_horiz_drill_kop.Kickoff_point_md_ouom, well_horiz_drill_kop.Kickoff_point_tvd, well_horiz_drill_kop.Kickoff_point_tvd_ouom, well_horiz_drill_kop.Lateral_hole_id, well_horiz_drill_kop.Node_id, well_horiz_drill_kop.Ppdm_guid, well_horiz_drill_kop.Remark, well_horiz_drill_kop.Row_changed_by, well_horiz_drill_kop.Row_changed_date, well_horiz_drill_kop.Row_created_by, well_horiz_drill_kop.Row_created_date, well_horiz_drill_kop.Row_effective_date, well_horiz_drill_kop.Row_expiry_date, well_horiz_drill_kop.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellHorizDrillKop(c *fiber.Ctx) error {
	var well_horiz_drill_kop dto.Well_horiz_drill_kop

	setDefaults(&well_horiz_drill_kop)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_kop); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_horiz_drill_kop.Uwi = id

    if well_horiz_drill_kop.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_horiz_drill_kop set source = :1, kickoff_point_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, kickoff_point_md = :6, kickoff_point_md_ouom = :7, kickoff_point_tvd = :8, kickoff_point_tvd_ouom = :9, lateral_hole_id = :10, node_id = :11, ppdm_guid = :12, remark = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_horiz_drill_kop.Row_changed_date = formatDate(well_horiz_drill_kop.Row_changed_date)
	well_horiz_drill_kop.Effective_date = formatDateString(well_horiz_drill_kop.Effective_date)
	well_horiz_drill_kop.Expiry_date = formatDateString(well_horiz_drill_kop.Expiry_date)
	well_horiz_drill_kop.Row_effective_date = formatDateString(well_horiz_drill_kop.Row_effective_date)
	well_horiz_drill_kop.Row_expiry_date = formatDateString(well_horiz_drill_kop.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_kop.Source, well_horiz_drill_kop.Kickoff_point_obs_no, well_horiz_drill_kop.Active_ind, well_horiz_drill_kop.Effective_date, well_horiz_drill_kop.Expiry_date, well_horiz_drill_kop.Kickoff_point_md, well_horiz_drill_kop.Kickoff_point_md_ouom, well_horiz_drill_kop.Kickoff_point_tvd, well_horiz_drill_kop.Kickoff_point_tvd_ouom, well_horiz_drill_kop.Lateral_hole_id, well_horiz_drill_kop.Node_id, well_horiz_drill_kop.Ppdm_guid, well_horiz_drill_kop.Remark, well_horiz_drill_kop.Row_changed_by, well_horiz_drill_kop.Row_changed_date, well_horiz_drill_kop.Row_created_by, well_horiz_drill_kop.Row_effective_date, well_horiz_drill_kop.Row_expiry_date, well_horiz_drill_kop.Row_quality, well_horiz_drill_kop.Uwi)
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

func PatchWellHorizDrillKop(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_horiz_drill_kop set "
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

func DeleteWellHorizDrillKop(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_horiz_drill_kop dto.Well_horiz_drill_kop
	well_horiz_drill_kop.Uwi = id

	stmt, err := db.Prepare("delete from well_horiz_drill_kop where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_horiz_drill_kop.Uwi)
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


