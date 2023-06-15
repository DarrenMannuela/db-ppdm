package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelNorthSea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_north_sea")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_north_sea

	for rows.Next() {
		var sp_parcel_north_sea dto.Sp_parcel_north_sea
		if err := rows.Scan(&sp_parcel_north_sea.Parcel_north_sea_id, &sp_parcel_north_sea.Active_ind, &sp_parcel_north_sea.Area_id, &sp_parcel_north_sea.Area_type, &sp_parcel_north_sea.Block_no, &sp_parcel_north_sea.Block_suffix, &sp_parcel_north_sea.Coord_system_id, &sp_parcel_north_sea.Description, &sp_parcel_north_sea.Effective_date, &sp_parcel_north_sea.Expiry_date, &sp_parcel_north_sea.Ppdm_guid, &sp_parcel_north_sea.Principal_meridian, &sp_parcel_north_sea.Quadrant, &sp_parcel_north_sea.Quadrant_prefix, &sp_parcel_north_sea.Reference_plan_num, &sp_parcel_north_sea.Remark, &sp_parcel_north_sea.Source, &sp_parcel_north_sea.Spatial_description_id, &sp_parcel_north_sea.Spatial_obs_no, &sp_parcel_north_sea.Subdivision, &sp_parcel_north_sea.Row_changed_by, &sp_parcel_north_sea.Row_changed_date, &sp_parcel_north_sea.Row_created_by, &sp_parcel_north_sea.Row_created_date, &sp_parcel_north_sea.Row_effective_date, &sp_parcel_north_sea.Row_expiry_date, &sp_parcel_north_sea.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_north_sea to result
		result = append(result, sp_parcel_north_sea)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelNorthSea(c *fiber.Ctx) error {
	var sp_parcel_north_sea dto.Sp_parcel_north_sea

	setDefaults(&sp_parcel_north_sea)

	if err := json.Unmarshal(c.Body(), &sp_parcel_north_sea); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_north_sea values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	sp_parcel_north_sea.Row_created_date = formatDate(sp_parcel_north_sea.Row_created_date)
	sp_parcel_north_sea.Row_changed_date = formatDate(sp_parcel_north_sea.Row_changed_date)
	sp_parcel_north_sea.Effective_date = formatDateString(sp_parcel_north_sea.Effective_date)
	sp_parcel_north_sea.Expiry_date = formatDateString(sp_parcel_north_sea.Expiry_date)
	sp_parcel_north_sea.Row_effective_date = formatDateString(sp_parcel_north_sea.Row_effective_date)
	sp_parcel_north_sea.Row_expiry_date = formatDateString(sp_parcel_north_sea.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_north_sea.Parcel_north_sea_id, sp_parcel_north_sea.Active_ind, sp_parcel_north_sea.Area_id, sp_parcel_north_sea.Area_type, sp_parcel_north_sea.Block_no, sp_parcel_north_sea.Block_suffix, sp_parcel_north_sea.Coord_system_id, sp_parcel_north_sea.Description, sp_parcel_north_sea.Effective_date, sp_parcel_north_sea.Expiry_date, sp_parcel_north_sea.Ppdm_guid, sp_parcel_north_sea.Principal_meridian, sp_parcel_north_sea.Quadrant, sp_parcel_north_sea.Quadrant_prefix, sp_parcel_north_sea.Reference_plan_num, sp_parcel_north_sea.Remark, sp_parcel_north_sea.Source, sp_parcel_north_sea.Spatial_description_id, sp_parcel_north_sea.Spatial_obs_no, sp_parcel_north_sea.Subdivision, sp_parcel_north_sea.Row_changed_by, sp_parcel_north_sea.Row_changed_date, sp_parcel_north_sea.Row_created_by, sp_parcel_north_sea.Row_created_date, sp_parcel_north_sea.Row_effective_date, sp_parcel_north_sea.Row_expiry_date, sp_parcel_north_sea.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelNorthSea(c *fiber.Ctx) error {
	var sp_parcel_north_sea dto.Sp_parcel_north_sea

	setDefaults(&sp_parcel_north_sea)

	if err := json.Unmarshal(c.Body(), &sp_parcel_north_sea); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_north_sea.Parcel_north_sea_id = id

    if sp_parcel_north_sea.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_north_sea set active_ind = :1, area_id = :2, area_type = :3, block_no = :4, block_suffix = :5, coord_system_id = :6, description = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, principal_meridian = :11, quadrant = :12, quadrant_prefix = :13, reference_plan_num = :14, remark = :15, source = :16, spatial_description_id = :17, spatial_obs_no = :18, subdivision = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where parcel_north_sea_id = :27")
	if err != nil {
		return err
	}

	sp_parcel_north_sea.Row_changed_date = formatDate(sp_parcel_north_sea.Row_changed_date)
	sp_parcel_north_sea.Effective_date = formatDateString(sp_parcel_north_sea.Effective_date)
	sp_parcel_north_sea.Expiry_date = formatDateString(sp_parcel_north_sea.Expiry_date)
	sp_parcel_north_sea.Row_effective_date = formatDateString(sp_parcel_north_sea.Row_effective_date)
	sp_parcel_north_sea.Row_expiry_date = formatDateString(sp_parcel_north_sea.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_north_sea.Active_ind, sp_parcel_north_sea.Area_id, sp_parcel_north_sea.Area_type, sp_parcel_north_sea.Block_no, sp_parcel_north_sea.Block_suffix, sp_parcel_north_sea.Coord_system_id, sp_parcel_north_sea.Description, sp_parcel_north_sea.Effective_date, sp_parcel_north_sea.Expiry_date, sp_parcel_north_sea.Ppdm_guid, sp_parcel_north_sea.Principal_meridian, sp_parcel_north_sea.Quadrant, sp_parcel_north_sea.Quadrant_prefix, sp_parcel_north_sea.Reference_plan_num, sp_parcel_north_sea.Remark, sp_parcel_north_sea.Source, sp_parcel_north_sea.Spatial_description_id, sp_parcel_north_sea.Spatial_obs_no, sp_parcel_north_sea.Subdivision, sp_parcel_north_sea.Row_changed_by, sp_parcel_north_sea.Row_changed_date, sp_parcel_north_sea.Row_created_by, sp_parcel_north_sea.Row_effective_date, sp_parcel_north_sea.Row_expiry_date, sp_parcel_north_sea.Row_quality, sp_parcel_north_sea.Parcel_north_sea_id)
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

func PatchSpParcelNorthSea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_north_sea set "
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
	query += " where parcel_north_sea_id = :id"

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

func DeleteSpParcelNorthSea(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_north_sea dto.Sp_parcel_north_sea
	sp_parcel_north_sea.Parcel_north_sea_id = id

	stmt, err := db.Prepare("delete from sp_parcel_north_sea where parcel_north_sea_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_north_sea.Parcel_north_sea_id)
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


