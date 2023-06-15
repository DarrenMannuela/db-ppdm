package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelCarter(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_carter")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_carter

	for rows.Next() {
		var sp_parcel_carter dto.Sp_parcel_carter
		if err := rows.Scan(&sp_parcel_carter.Parcel_carter_id, &sp_parcel_carter.Active_ind, &sp_parcel_carter.Area_id, &sp_parcel_carter.Area_type, &sp_parcel_carter.Carter_range, &sp_parcel_carter.Carter_section, &sp_parcel_carter.Carter_township, &sp_parcel_carter.Coord_system_id, &sp_parcel_carter.Description, &sp_parcel_carter.Effective_date, &sp_parcel_carter.Ew_direction, &sp_parcel_carter.Expiry_date, &sp_parcel_carter.Map_quad_min, &sp_parcel_carter.Map_quad_name, &sp_parcel_carter.Ns_direction, &sp_parcel_carter.Ppdm_guid, &sp_parcel_carter.Reference_plan_num, &sp_parcel_carter.Remark, &sp_parcel_carter.Source, &sp_parcel_carter.Spatial_description_id, &sp_parcel_carter.Spatial_obs_no, &sp_parcel_carter.Spot_code, &sp_parcel_carter.Row_changed_by, &sp_parcel_carter.Row_changed_date, &sp_parcel_carter.Row_created_by, &sp_parcel_carter.Row_created_date, &sp_parcel_carter.Row_effective_date, &sp_parcel_carter.Row_expiry_date, &sp_parcel_carter.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_carter to result
		result = append(result, sp_parcel_carter)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelCarter(c *fiber.Ctx) error {
	var sp_parcel_carter dto.Sp_parcel_carter

	setDefaults(&sp_parcel_carter)

	if err := json.Unmarshal(c.Body(), &sp_parcel_carter); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_carter values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	sp_parcel_carter.Row_created_date = formatDate(sp_parcel_carter.Row_created_date)
	sp_parcel_carter.Row_changed_date = formatDate(sp_parcel_carter.Row_changed_date)
	sp_parcel_carter.Effective_date = formatDateString(sp_parcel_carter.Effective_date)
	sp_parcel_carter.Expiry_date = formatDateString(sp_parcel_carter.Expiry_date)
	sp_parcel_carter.Row_effective_date = formatDateString(sp_parcel_carter.Row_effective_date)
	sp_parcel_carter.Row_expiry_date = formatDateString(sp_parcel_carter.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_carter.Parcel_carter_id, sp_parcel_carter.Active_ind, sp_parcel_carter.Area_id, sp_parcel_carter.Area_type, sp_parcel_carter.Carter_range, sp_parcel_carter.Carter_section, sp_parcel_carter.Carter_township, sp_parcel_carter.Coord_system_id, sp_parcel_carter.Description, sp_parcel_carter.Effective_date, sp_parcel_carter.Ew_direction, sp_parcel_carter.Expiry_date, sp_parcel_carter.Map_quad_min, sp_parcel_carter.Map_quad_name, sp_parcel_carter.Ns_direction, sp_parcel_carter.Ppdm_guid, sp_parcel_carter.Reference_plan_num, sp_parcel_carter.Remark, sp_parcel_carter.Source, sp_parcel_carter.Spatial_description_id, sp_parcel_carter.Spatial_obs_no, sp_parcel_carter.Spot_code, sp_parcel_carter.Row_changed_by, sp_parcel_carter.Row_changed_date, sp_parcel_carter.Row_created_by, sp_parcel_carter.Row_created_date, sp_parcel_carter.Row_effective_date, sp_parcel_carter.Row_expiry_date, sp_parcel_carter.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelCarter(c *fiber.Ctx) error {
	var sp_parcel_carter dto.Sp_parcel_carter

	setDefaults(&sp_parcel_carter)

	if err := json.Unmarshal(c.Body(), &sp_parcel_carter); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_carter.Parcel_carter_id = id

    if sp_parcel_carter.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_carter set active_ind = :1, area_id = :2, area_type = :3, carter_range = :4, carter_section = :5, carter_township = :6, coord_system_id = :7, description = :8, effective_date = :9, ew_direction = :10, expiry_date = :11, map_quad_min = :12, map_quad_name = :13, ns_direction = :14, ppdm_guid = :15, reference_plan_num = :16, remark = :17, source = :18, spatial_description_id = :19, spatial_obs_no = :20, spot_code = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where parcel_carter_id = :29")
	if err != nil {
		return err
	}

	sp_parcel_carter.Row_changed_date = formatDate(sp_parcel_carter.Row_changed_date)
	sp_parcel_carter.Effective_date = formatDateString(sp_parcel_carter.Effective_date)
	sp_parcel_carter.Expiry_date = formatDateString(sp_parcel_carter.Expiry_date)
	sp_parcel_carter.Row_effective_date = formatDateString(sp_parcel_carter.Row_effective_date)
	sp_parcel_carter.Row_expiry_date = formatDateString(sp_parcel_carter.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_carter.Active_ind, sp_parcel_carter.Area_id, sp_parcel_carter.Area_type, sp_parcel_carter.Carter_range, sp_parcel_carter.Carter_section, sp_parcel_carter.Carter_township, sp_parcel_carter.Coord_system_id, sp_parcel_carter.Description, sp_parcel_carter.Effective_date, sp_parcel_carter.Ew_direction, sp_parcel_carter.Expiry_date, sp_parcel_carter.Map_quad_min, sp_parcel_carter.Map_quad_name, sp_parcel_carter.Ns_direction, sp_parcel_carter.Ppdm_guid, sp_parcel_carter.Reference_plan_num, sp_parcel_carter.Remark, sp_parcel_carter.Source, sp_parcel_carter.Spatial_description_id, sp_parcel_carter.Spatial_obs_no, sp_parcel_carter.Spot_code, sp_parcel_carter.Row_changed_by, sp_parcel_carter.Row_changed_date, sp_parcel_carter.Row_created_by, sp_parcel_carter.Row_effective_date, sp_parcel_carter.Row_expiry_date, sp_parcel_carter.Row_quality, sp_parcel_carter.Parcel_carter_id)
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

func PatchSpParcelCarter(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_carter set "
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
	query += " where parcel_carter_id = :id"

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

func DeleteSpParcelCarter(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_carter dto.Sp_parcel_carter
	sp_parcel_carter.Parcel_carter_id = id

	stmt, err := db.Prepare("delete from sp_parcel_carter where parcel_carter_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_carter.Parcel_carter_id)
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


