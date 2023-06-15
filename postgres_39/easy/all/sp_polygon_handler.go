package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpPolygon(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_polygon")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_polygon

	for rows.Next() {
		var sp_polygon dto.Sp_polygon
		if err := rows.Scan(&sp_polygon.Polygon_id, &sp_polygon.Acquisition_id, &sp_polygon.Active_ind, &sp_polygon.Boundary_direction, &sp_polygon.Contained_by_polygon_id, &sp_polygon.Coord_system_id, &sp_polygon.Datum_elev, &sp_polygon.Datum_elev_ouom, &sp_polygon.Effective_date, &sp_polygon.Exclusion_ind, &sp_polygon.Expiry_date, &sp_polygon.Inclusion_ind, &sp_polygon.Local_coord_system_id, &sp_polygon.Location_type, &sp_polygon.Max_plot_scale, &sp_polygon.Min_plot_scale, &sp_polygon.Polygon_set_id, &sp_polygon.Ppdm_guid, &sp_polygon.Preferred_ind, &sp_polygon.Reference_datum, &sp_polygon.Remark, &sp_polygon.Source, &sp_polygon.Spatial_description_id, &sp_polygon.Spatial_obs_no, &sp_polygon.Row_changed_by, &sp_polygon.Row_changed_date, &sp_polygon.Row_created_by, &sp_polygon.Row_created_date, &sp_polygon.Row_effective_date, &sp_polygon.Row_expiry_date, &sp_polygon.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_polygon to result
		result = append(result, sp_polygon)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpPolygon(c *fiber.Ctx) error {
	var sp_polygon dto.Sp_polygon

	setDefaults(&sp_polygon)

	if err := json.Unmarshal(c.Body(), &sp_polygon); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_polygon values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	sp_polygon.Row_created_date = formatDate(sp_polygon.Row_created_date)
	sp_polygon.Row_changed_date = formatDate(sp_polygon.Row_changed_date)
	sp_polygon.Effective_date = formatDateString(sp_polygon.Effective_date)
	sp_polygon.Expiry_date = formatDateString(sp_polygon.Expiry_date)
	sp_polygon.Row_effective_date = formatDateString(sp_polygon.Row_effective_date)
	sp_polygon.Row_expiry_date = formatDateString(sp_polygon.Row_expiry_date)






	rows, err := stmt.Exec(sp_polygon.Polygon_id, sp_polygon.Acquisition_id, sp_polygon.Active_ind, sp_polygon.Boundary_direction, sp_polygon.Contained_by_polygon_id, sp_polygon.Coord_system_id, sp_polygon.Datum_elev, sp_polygon.Datum_elev_ouom, sp_polygon.Effective_date, sp_polygon.Exclusion_ind, sp_polygon.Expiry_date, sp_polygon.Inclusion_ind, sp_polygon.Local_coord_system_id, sp_polygon.Location_type, sp_polygon.Max_plot_scale, sp_polygon.Min_plot_scale, sp_polygon.Polygon_set_id, sp_polygon.Ppdm_guid, sp_polygon.Preferred_ind, sp_polygon.Reference_datum, sp_polygon.Remark, sp_polygon.Source, sp_polygon.Spatial_description_id, sp_polygon.Spatial_obs_no, sp_polygon.Row_changed_by, sp_polygon.Row_changed_date, sp_polygon.Row_created_by, sp_polygon.Row_created_date, sp_polygon.Row_effective_date, sp_polygon.Row_expiry_date, sp_polygon.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpPolygon(c *fiber.Ctx) error {
	var sp_polygon dto.Sp_polygon

	setDefaults(&sp_polygon)

	if err := json.Unmarshal(c.Body(), &sp_polygon); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_polygon.Polygon_id = id

    if sp_polygon.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_polygon set acquisition_id = :1, active_ind = :2, boundary_direction = :3, contained_by_polygon_id = :4, coord_system_id = :5, datum_elev = :6, datum_elev_ouom = :7, effective_date = :8, exclusion_ind = :9, expiry_date = :10, inclusion_ind = :11, local_coord_system_id = :12, location_type = :13, max_plot_scale = :14, min_plot_scale = :15, polygon_set_id = :16, ppdm_guid = :17, preferred_ind = :18, reference_datum = :19, remark = :20, source = :21, spatial_description_id = :22, spatial_obs_no = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where polygon_id = :31")
	if err != nil {
		return err
	}

	sp_polygon.Row_changed_date = formatDate(sp_polygon.Row_changed_date)
	sp_polygon.Effective_date = formatDateString(sp_polygon.Effective_date)
	sp_polygon.Expiry_date = formatDateString(sp_polygon.Expiry_date)
	sp_polygon.Row_effective_date = formatDateString(sp_polygon.Row_effective_date)
	sp_polygon.Row_expiry_date = formatDateString(sp_polygon.Row_expiry_date)






	rows, err := stmt.Exec(sp_polygon.Acquisition_id, sp_polygon.Active_ind, sp_polygon.Boundary_direction, sp_polygon.Contained_by_polygon_id, sp_polygon.Coord_system_id, sp_polygon.Datum_elev, sp_polygon.Datum_elev_ouom, sp_polygon.Effective_date, sp_polygon.Exclusion_ind, sp_polygon.Expiry_date, sp_polygon.Inclusion_ind, sp_polygon.Local_coord_system_id, sp_polygon.Location_type, sp_polygon.Max_plot_scale, sp_polygon.Min_plot_scale, sp_polygon.Polygon_set_id, sp_polygon.Ppdm_guid, sp_polygon.Preferred_ind, sp_polygon.Reference_datum, sp_polygon.Remark, sp_polygon.Source, sp_polygon.Spatial_description_id, sp_polygon.Spatial_obs_no, sp_polygon.Row_changed_by, sp_polygon.Row_changed_date, sp_polygon.Row_created_by, sp_polygon.Row_effective_date, sp_polygon.Row_expiry_date, sp_polygon.Row_quality, sp_polygon.Polygon_id)
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

func PatchSpPolygon(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_polygon set "
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
	query += " where polygon_id = :id"

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

func DeleteSpPolygon(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_polygon dto.Sp_polygon
	sp_polygon.Polygon_id = id

	stmt, err := db.Prepare("delete from sp_polygon where polygon_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_polygon.Polygon_id)
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


