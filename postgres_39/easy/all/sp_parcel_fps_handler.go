package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelFps(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_fps")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_fps

	for rows.Next() {
		var sp_parcel_fps dto.Sp_parcel_fps
		if err := rows.Scan(&sp_parcel_fps.Parcel_fps_id, &sp_parcel_fps.Active_ind, &sp_parcel_fps.Area_id, &sp_parcel_fps.Area_type, &sp_parcel_fps.Coord_system_id, &sp_parcel_fps.Description, &sp_parcel_fps.Effective_date, &sp_parcel_fps.Expiry_date, &sp_parcel_fps.Grid, &sp_parcel_fps.Ppdm_guid, &sp_parcel_fps.Reference_latitude, &sp_parcel_fps.Reference_longitude, &sp_parcel_fps.Reference_plan_num, &sp_parcel_fps.Remark, &sp_parcel_fps.Section, &sp_parcel_fps.Source, &sp_parcel_fps.Spatial_description_id, &sp_parcel_fps.Spatial_obs_no, &sp_parcel_fps.Unit, &sp_parcel_fps.Row_changed_by, &sp_parcel_fps.Row_changed_date, &sp_parcel_fps.Row_created_by, &sp_parcel_fps.Row_created_date, &sp_parcel_fps.Row_effective_date, &sp_parcel_fps.Row_expiry_date, &sp_parcel_fps.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_fps to result
		result = append(result, sp_parcel_fps)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelFps(c *fiber.Ctx) error {
	var sp_parcel_fps dto.Sp_parcel_fps

	setDefaults(&sp_parcel_fps)

	if err := json.Unmarshal(c.Body(), &sp_parcel_fps); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_fps values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	sp_parcel_fps.Row_created_date = formatDate(sp_parcel_fps.Row_created_date)
	sp_parcel_fps.Row_changed_date = formatDate(sp_parcel_fps.Row_changed_date)
	sp_parcel_fps.Effective_date = formatDateString(sp_parcel_fps.Effective_date)
	sp_parcel_fps.Expiry_date = formatDateString(sp_parcel_fps.Expiry_date)
	sp_parcel_fps.Row_effective_date = formatDateString(sp_parcel_fps.Row_effective_date)
	sp_parcel_fps.Row_expiry_date = formatDateString(sp_parcel_fps.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_fps.Parcel_fps_id, sp_parcel_fps.Active_ind, sp_parcel_fps.Area_id, sp_parcel_fps.Area_type, sp_parcel_fps.Coord_system_id, sp_parcel_fps.Description, sp_parcel_fps.Effective_date, sp_parcel_fps.Expiry_date, sp_parcel_fps.Grid, sp_parcel_fps.Ppdm_guid, sp_parcel_fps.Reference_latitude, sp_parcel_fps.Reference_longitude, sp_parcel_fps.Reference_plan_num, sp_parcel_fps.Remark, sp_parcel_fps.Section, sp_parcel_fps.Source, sp_parcel_fps.Spatial_description_id, sp_parcel_fps.Spatial_obs_no, sp_parcel_fps.Unit, sp_parcel_fps.Row_changed_by, sp_parcel_fps.Row_changed_date, sp_parcel_fps.Row_created_by, sp_parcel_fps.Row_created_date, sp_parcel_fps.Row_effective_date, sp_parcel_fps.Row_expiry_date, sp_parcel_fps.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelFps(c *fiber.Ctx) error {
	var sp_parcel_fps dto.Sp_parcel_fps

	setDefaults(&sp_parcel_fps)

	if err := json.Unmarshal(c.Body(), &sp_parcel_fps); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_fps.Parcel_fps_id = id

    if sp_parcel_fps.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_fps set active_ind = :1, area_id = :2, area_type = :3, coord_system_id = :4, description = :5, effective_date = :6, expiry_date = :7, grid = :8, ppdm_guid = :9, reference_latitude = :10, reference_longitude = :11, reference_plan_num = :12, remark = :13, section = :14, source = :15, spatial_description_id = :16, spatial_obs_no = :17, unit = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where parcel_fps_id = :26")
	if err != nil {
		return err
	}

	sp_parcel_fps.Row_changed_date = formatDate(sp_parcel_fps.Row_changed_date)
	sp_parcel_fps.Effective_date = formatDateString(sp_parcel_fps.Effective_date)
	sp_parcel_fps.Expiry_date = formatDateString(sp_parcel_fps.Expiry_date)
	sp_parcel_fps.Row_effective_date = formatDateString(sp_parcel_fps.Row_effective_date)
	sp_parcel_fps.Row_expiry_date = formatDateString(sp_parcel_fps.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_fps.Active_ind, sp_parcel_fps.Area_id, sp_parcel_fps.Area_type, sp_parcel_fps.Coord_system_id, sp_parcel_fps.Description, sp_parcel_fps.Effective_date, sp_parcel_fps.Expiry_date, sp_parcel_fps.Grid, sp_parcel_fps.Ppdm_guid, sp_parcel_fps.Reference_latitude, sp_parcel_fps.Reference_longitude, sp_parcel_fps.Reference_plan_num, sp_parcel_fps.Remark, sp_parcel_fps.Section, sp_parcel_fps.Source, sp_parcel_fps.Spatial_description_id, sp_parcel_fps.Spatial_obs_no, sp_parcel_fps.Unit, sp_parcel_fps.Row_changed_by, sp_parcel_fps.Row_changed_date, sp_parcel_fps.Row_created_by, sp_parcel_fps.Row_effective_date, sp_parcel_fps.Row_expiry_date, sp_parcel_fps.Row_quality, sp_parcel_fps.Parcel_fps_id)
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

func PatchSpParcelFps(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_fps set "
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
	query += " where parcel_fps_id = :id"

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

func DeleteSpParcelFps(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_fps dto.Sp_parcel_fps
	sp_parcel_fps.Parcel_fps_id = id

	stmt, err := db.Prepare("delete from sp_parcel_fps where parcel_fps_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_fps.Parcel_fps_id)
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


