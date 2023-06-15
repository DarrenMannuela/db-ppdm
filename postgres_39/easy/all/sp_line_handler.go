package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpLine(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_line")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_line

	for rows.Next() {
		var sp_line dto.Sp_line
		if err := rows.Scan(&sp_line.Line_id, &sp_line.Acquisition_id, &sp_line.Active_ind, &sp_line.Coord_system_id, &sp_line.Datum_elev, &sp_line.Datum_elev_ouom, &sp_line.Effective_date, &sp_line.Expiry_date, &sp_line.Line_set_id, &sp_line.Local_coord_system_id, &sp_line.Location_type, &sp_line.Max_plot_scale, &sp_line.Min_plot_scale, &sp_line.Ppdm_guid, &sp_line.Preferred_ind, &sp_line.Reference_datum, &sp_line.Remark, &sp_line.Source, &sp_line.Spatial_description_id, &sp_line.Spatial_obs_no, &sp_line.Row_changed_by, &sp_line.Row_changed_date, &sp_line.Row_created_by, &sp_line.Row_created_date, &sp_line.Row_effective_date, &sp_line.Row_expiry_date, &sp_line.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_line to result
		result = append(result, sp_line)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpLine(c *fiber.Ctx) error {
	var sp_line dto.Sp_line

	setDefaults(&sp_line)

	if err := json.Unmarshal(c.Body(), &sp_line); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_line values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	sp_line.Row_created_date = formatDate(sp_line.Row_created_date)
	sp_line.Row_changed_date = formatDate(sp_line.Row_changed_date)
	sp_line.Effective_date = formatDateString(sp_line.Effective_date)
	sp_line.Expiry_date = formatDateString(sp_line.Expiry_date)
	sp_line.Row_effective_date = formatDateString(sp_line.Row_effective_date)
	sp_line.Row_expiry_date = formatDateString(sp_line.Row_expiry_date)






	rows, err := stmt.Exec(sp_line.Line_id, sp_line.Acquisition_id, sp_line.Active_ind, sp_line.Coord_system_id, sp_line.Datum_elev, sp_line.Datum_elev_ouom, sp_line.Effective_date, sp_line.Expiry_date, sp_line.Line_set_id, sp_line.Local_coord_system_id, sp_line.Location_type, sp_line.Max_plot_scale, sp_line.Min_plot_scale, sp_line.Ppdm_guid, sp_line.Preferred_ind, sp_line.Reference_datum, sp_line.Remark, sp_line.Source, sp_line.Spatial_description_id, sp_line.Spatial_obs_no, sp_line.Row_changed_by, sp_line.Row_changed_date, sp_line.Row_created_by, sp_line.Row_created_date, sp_line.Row_effective_date, sp_line.Row_expiry_date, sp_line.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpLine(c *fiber.Ctx) error {
	var sp_line dto.Sp_line

	setDefaults(&sp_line)

	if err := json.Unmarshal(c.Body(), &sp_line); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_line.Line_id = id

    if sp_line.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_line set acquisition_id = :1, active_ind = :2, coord_system_id = :3, datum_elev = :4, datum_elev_ouom = :5, effective_date = :6, expiry_date = :7, line_set_id = :8, local_coord_system_id = :9, location_type = :10, max_plot_scale = :11, min_plot_scale = :12, ppdm_guid = :13, preferred_ind = :14, reference_datum = :15, remark = :16, source = :17, spatial_description_id = :18, spatial_obs_no = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where line_id = :27")
	if err != nil {
		return err
	}

	sp_line.Row_changed_date = formatDate(sp_line.Row_changed_date)
	sp_line.Effective_date = formatDateString(sp_line.Effective_date)
	sp_line.Expiry_date = formatDateString(sp_line.Expiry_date)
	sp_line.Row_effective_date = formatDateString(sp_line.Row_effective_date)
	sp_line.Row_expiry_date = formatDateString(sp_line.Row_expiry_date)






	rows, err := stmt.Exec(sp_line.Acquisition_id, sp_line.Active_ind, sp_line.Coord_system_id, sp_line.Datum_elev, sp_line.Datum_elev_ouom, sp_line.Effective_date, sp_line.Expiry_date, sp_line.Line_set_id, sp_line.Local_coord_system_id, sp_line.Location_type, sp_line.Max_plot_scale, sp_line.Min_plot_scale, sp_line.Ppdm_guid, sp_line.Preferred_ind, sp_line.Reference_datum, sp_line.Remark, sp_line.Source, sp_line.Spatial_description_id, sp_line.Spatial_obs_no, sp_line.Row_changed_by, sp_line.Row_changed_date, sp_line.Row_created_by, sp_line.Row_effective_date, sp_line.Row_expiry_date, sp_line.Row_quality, sp_line.Line_id)
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

func PatchSpLine(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_line set "
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
	query += " where line_id = :id"

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

func DeleteSpLine(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_line dto.Sp_line
	sp_line.Line_id = id

	stmt, err := db.Prepare("delete from sp_line where line_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_line.Line_id)
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


