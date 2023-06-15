package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelOffshore(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_offshore")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_offshore

	for rows.Next() {
		var sp_parcel_offshore dto.Sp_parcel_offshore
		if err := rows.Scan(&sp_parcel_offshore.Parcel_offshore_id, &sp_parcel_offshore.Active_ind, &sp_parcel_offshore.Area_id, &sp_parcel_offshore.Area_type, &sp_parcel_offshore.Block_addition, &sp_parcel_offshore.Block_modifier, &sp_parcel_offshore.Coord_system_id, &sp_parcel_offshore.Description, &sp_parcel_offshore.Effective_date, &sp_parcel_offshore.Ew_direction, &sp_parcel_offshore.Expiry_date, &sp_parcel_offshore.Grid_block_range, &sp_parcel_offshore.Grid_block_tier, &sp_parcel_offshore.Ns_direction, &sp_parcel_offshore.Ocs_num, &sp_parcel_offshore.Offshore_area_code, &sp_parcel_offshore.Offshore_block, &sp_parcel_offshore.Ppdm_guid, &sp_parcel_offshore.Reference_plan_num, &sp_parcel_offshore.Remark, &sp_parcel_offshore.Source, &sp_parcel_offshore.Spatial_description_id, &sp_parcel_offshore.Spatial_obs_no, &sp_parcel_offshore.Utm_quadrant, &sp_parcel_offshore.Water_bottom_zone, &sp_parcel_offshore.Row_changed_by, &sp_parcel_offshore.Row_changed_date, &sp_parcel_offshore.Row_created_by, &sp_parcel_offshore.Row_created_date, &sp_parcel_offshore.Row_effective_date, &sp_parcel_offshore.Row_expiry_date, &sp_parcel_offshore.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_offshore to result
		result = append(result, sp_parcel_offshore)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelOffshore(c *fiber.Ctx) error {
	var sp_parcel_offshore dto.Sp_parcel_offshore

	setDefaults(&sp_parcel_offshore)

	if err := json.Unmarshal(c.Body(), &sp_parcel_offshore); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_offshore values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	sp_parcel_offshore.Row_created_date = formatDate(sp_parcel_offshore.Row_created_date)
	sp_parcel_offshore.Row_changed_date = formatDate(sp_parcel_offshore.Row_changed_date)
	sp_parcel_offshore.Effective_date = formatDateString(sp_parcel_offshore.Effective_date)
	sp_parcel_offshore.Expiry_date = formatDateString(sp_parcel_offshore.Expiry_date)
	sp_parcel_offshore.Row_effective_date = formatDateString(sp_parcel_offshore.Row_effective_date)
	sp_parcel_offshore.Row_expiry_date = formatDateString(sp_parcel_offshore.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_offshore.Parcel_offshore_id, sp_parcel_offshore.Active_ind, sp_parcel_offshore.Area_id, sp_parcel_offshore.Area_type, sp_parcel_offshore.Block_addition, sp_parcel_offshore.Block_modifier, sp_parcel_offshore.Coord_system_id, sp_parcel_offshore.Description, sp_parcel_offshore.Effective_date, sp_parcel_offshore.Ew_direction, sp_parcel_offshore.Expiry_date, sp_parcel_offshore.Grid_block_range, sp_parcel_offshore.Grid_block_tier, sp_parcel_offshore.Ns_direction, sp_parcel_offshore.Ocs_num, sp_parcel_offshore.Offshore_area_code, sp_parcel_offshore.Offshore_block, sp_parcel_offshore.Ppdm_guid, sp_parcel_offshore.Reference_plan_num, sp_parcel_offshore.Remark, sp_parcel_offshore.Source, sp_parcel_offshore.Spatial_description_id, sp_parcel_offshore.Spatial_obs_no, sp_parcel_offshore.Utm_quadrant, sp_parcel_offshore.Water_bottom_zone, sp_parcel_offshore.Row_changed_by, sp_parcel_offshore.Row_changed_date, sp_parcel_offshore.Row_created_by, sp_parcel_offshore.Row_created_date, sp_parcel_offshore.Row_effective_date, sp_parcel_offshore.Row_expiry_date, sp_parcel_offshore.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelOffshore(c *fiber.Ctx) error {
	var sp_parcel_offshore dto.Sp_parcel_offshore

	setDefaults(&sp_parcel_offshore)

	if err := json.Unmarshal(c.Body(), &sp_parcel_offshore); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_offshore.Parcel_offshore_id = id

    if sp_parcel_offshore.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_offshore set active_ind = :1, area_id = :2, area_type = :3, block_addition = :4, block_modifier = :5, coord_system_id = :6, description = :7, effective_date = :8, ew_direction = :9, expiry_date = :10, grid_block_range = :11, grid_block_tier = :12, ns_direction = :13, ocs_num = :14, offshore_area_code = :15, offshore_block = :16, ppdm_guid = :17, reference_plan_num = :18, remark = :19, source = :20, spatial_description_id = :21, spatial_obs_no = :22, utm_quadrant = :23, water_bottom_zone = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where parcel_offshore_id = :32")
	if err != nil {
		return err
	}

	sp_parcel_offshore.Row_changed_date = formatDate(sp_parcel_offshore.Row_changed_date)
	sp_parcel_offshore.Effective_date = formatDateString(sp_parcel_offshore.Effective_date)
	sp_parcel_offshore.Expiry_date = formatDateString(sp_parcel_offshore.Expiry_date)
	sp_parcel_offshore.Row_effective_date = formatDateString(sp_parcel_offshore.Row_effective_date)
	sp_parcel_offshore.Row_expiry_date = formatDateString(sp_parcel_offshore.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_offshore.Active_ind, sp_parcel_offshore.Area_id, sp_parcel_offshore.Area_type, sp_parcel_offshore.Block_addition, sp_parcel_offshore.Block_modifier, sp_parcel_offshore.Coord_system_id, sp_parcel_offshore.Description, sp_parcel_offshore.Effective_date, sp_parcel_offshore.Ew_direction, sp_parcel_offshore.Expiry_date, sp_parcel_offshore.Grid_block_range, sp_parcel_offshore.Grid_block_tier, sp_parcel_offshore.Ns_direction, sp_parcel_offshore.Ocs_num, sp_parcel_offshore.Offshore_area_code, sp_parcel_offshore.Offshore_block, sp_parcel_offshore.Ppdm_guid, sp_parcel_offshore.Reference_plan_num, sp_parcel_offshore.Remark, sp_parcel_offshore.Source, sp_parcel_offshore.Spatial_description_id, sp_parcel_offshore.Spatial_obs_no, sp_parcel_offshore.Utm_quadrant, sp_parcel_offshore.Water_bottom_zone, sp_parcel_offshore.Row_changed_by, sp_parcel_offshore.Row_changed_date, sp_parcel_offshore.Row_created_by, sp_parcel_offshore.Row_effective_date, sp_parcel_offshore.Row_expiry_date, sp_parcel_offshore.Row_quality, sp_parcel_offshore.Parcel_offshore_id)
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

func PatchSpParcelOffshore(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_offshore set "
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
	query += " where parcel_offshore_id = :id"

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

func DeleteSpParcelOffshore(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_offshore dto.Sp_parcel_offshore
	sp_parcel_offshore.Parcel_offshore_id = id

	stmt, err := db.Prepare("delete from sp_parcel_offshore where parcel_offshore_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_offshore.Parcel_offshore_id)
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


