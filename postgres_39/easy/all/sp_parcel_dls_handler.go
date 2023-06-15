package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelDls(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_dls")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_dls

	for rows.Next() {
		var sp_parcel_dls dto.Sp_parcel_dls
		if err := rows.Scan(&sp_parcel_dls.Parcel_dls_id, &sp_parcel_dls.Active_ind, &sp_parcel_dls.Area_id, &sp_parcel_dls.Area_type, &sp_parcel_dls.Coord_system_id, &sp_parcel_dls.Description, &sp_parcel_dls.Dls_legal_subdivision, &sp_parcel_dls.Dls_meridian, &sp_parcel_dls.Dls_meridian_direction, &sp_parcel_dls.Dls_quarter_section, &sp_parcel_dls.Dls_quarter_section_quarter, &sp_parcel_dls.Dls_range, &sp_parcel_dls.Dls_range_modifier, &sp_parcel_dls.Dls_section, &sp_parcel_dls.Dls_township, &sp_parcel_dls.Dls_township_modifier, &sp_parcel_dls.Effective_date, &sp_parcel_dls.Expiry_date, &sp_parcel_dls.Local_coord_system_id, &sp_parcel_dls.Ppdm_guid, &sp_parcel_dls.Principal_meridian, &sp_parcel_dls.Reference_plan_num, &sp_parcel_dls.Remark, &sp_parcel_dls.Source, &sp_parcel_dls.Spatial_description_id, &sp_parcel_dls.Spatial_obs_no, &sp_parcel_dls.Row_changed_by, &sp_parcel_dls.Row_changed_date, &sp_parcel_dls.Row_created_by, &sp_parcel_dls.Row_created_date, &sp_parcel_dls.Row_effective_date, &sp_parcel_dls.Row_expiry_date, &sp_parcel_dls.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_dls to result
		result = append(result, sp_parcel_dls)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelDls(c *fiber.Ctx) error {
	var sp_parcel_dls dto.Sp_parcel_dls

	setDefaults(&sp_parcel_dls)

	if err := json.Unmarshal(c.Body(), &sp_parcel_dls); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_dls values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	sp_parcel_dls.Row_created_date = formatDate(sp_parcel_dls.Row_created_date)
	sp_parcel_dls.Row_changed_date = formatDate(sp_parcel_dls.Row_changed_date)
	sp_parcel_dls.Effective_date = formatDateString(sp_parcel_dls.Effective_date)
	sp_parcel_dls.Expiry_date = formatDateString(sp_parcel_dls.Expiry_date)
	sp_parcel_dls.Row_effective_date = formatDateString(sp_parcel_dls.Row_effective_date)
	sp_parcel_dls.Row_expiry_date = formatDateString(sp_parcel_dls.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_dls.Parcel_dls_id, sp_parcel_dls.Active_ind, sp_parcel_dls.Area_id, sp_parcel_dls.Area_type, sp_parcel_dls.Coord_system_id, sp_parcel_dls.Description, sp_parcel_dls.Dls_legal_subdivision, sp_parcel_dls.Dls_meridian, sp_parcel_dls.Dls_meridian_direction, sp_parcel_dls.Dls_quarter_section, sp_parcel_dls.Dls_quarter_section_quarter, sp_parcel_dls.Dls_range, sp_parcel_dls.Dls_range_modifier, sp_parcel_dls.Dls_section, sp_parcel_dls.Dls_township, sp_parcel_dls.Dls_township_modifier, sp_parcel_dls.Effective_date, sp_parcel_dls.Expiry_date, sp_parcel_dls.Local_coord_system_id, sp_parcel_dls.Ppdm_guid, sp_parcel_dls.Principal_meridian, sp_parcel_dls.Reference_plan_num, sp_parcel_dls.Remark, sp_parcel_dls.Source, sp_parcel_dls.Spatial_description_id, sp_parcel_dls.Spatial_obs_no, sp_parcel_dls.Row_changed_by, sp_parcel_dls.Row_changed_date, sp_parcel_dls.Row_created_by, sp_parcel_dls.Row_created_date, sp_parcel_dls.Row_effective_date, sp_parcel_dls.Row_expiry_date, sp_parcel_dls.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelDls(c *fiber.Ctx) error {
	var sp_parcel_dls dto.Sp_parcel_dls

	setDefaults(&sp_parcel_dls)

	if err := json.Unmarshal(c.Body(), &sp_parcel_dls); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_dls.Parcel_dls_id = id

    if sp_parcel_dls.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_dls set active_ind = :1, area_id = :2, area_type = :3, coord_system_id = :4, description = :5, dls_legal_subdivision = :6, dls_meridian = :7, dls_meridian_direction = :8, dls_quarter_section = :9, dls_quarter_section_quarter = :10, dls_range = :11, dls_range_modifier = :12, dls_section = :13, dls_township = :14, dls_township_modifier = :15, effective_date = :16, expiry_date = :17, local_coord_system_id = :18, ppdm_guid = :19, principal_meridian = :20, reference_plan_num = :21, remark = :22, source = :23, spatial_description_id = :24, spatial_obs_no = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where parcel_dls_id = :33")
	if err != nil {
		return err
	}

	sp_parcel_dls.Row_changed_date = formatDate(sp_parcel_dls.Row_changed_date)
	sp_parcel_dls.Effective_date = formatDateString(sp_parcel_dls.Effective_date)
	sp_parcel_dls.Expiry_date = formatDateString(sp_parcel_dls.Expiry_date)
	sp_parcel_dls.Row_effective_date = formatDateString(sp_parcel_dls.Row_effective_date)
	sp_parcel_dls.Row_expiry_date = formatDateString(sp_parcel_dls.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_dls.Active_ind, sp_parcel_dls.Area_id, sp_parcel_dls.Area_type, sp_parcel_dls.Coord_system_id, sp_parcel_dls.Description, sp_parcel_dls.Dls_legal_subdivision, sp_parcel_dls.Dls_meridian, sp_parcel_dls.Dls_meridian_direction, sp_parcel_dls.Dls_quarter_section, sp_parcel_dls.Dls_quarter_section_quarter, sp_parcel_dls.Dls_range, sp_parcel_dls.Dls_range_modifier, sp_parcel_dls.Dls_section, sp_parcel_dls.Dls_township, sp_parcel_dls.Dls_township_modifier, sp_parcel_dls.Effective_date, sp_parcel_dls.Expiry_date, sp_parcel_dls.Local_coord_system_id, sp_parcel_dls.Ppdm_guid, sp_parcel_dls.Principal_meridian, sp_parcel_dls.Reference_plan_num, sp_parcel_dls.Remark, sp_parcel_dls.Source, sp_parcel_dls.Spatial_description_id, sp_parcel_dls.Spatial_obs_no, sp_parcel_dls.Row_changed_by, sp_parcel_dls.Row_changed_date, sp_parcel_dls.Row_created_by, sp_parcel_dls.Row_effective_date, sp_parcel_dls.Row_expiry_date, sp_parcel_dls.Row_quality, sp_parcel_dls.Parcel_dls_id)
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

func PatchSpParcelDls(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_dls set "
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
	query += " where parcel_dls_id = :id"

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

func DeleteSpParcelDls(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_dls dto.Sp_parcel_dls
	sp_parcel_dls.Parcel_dls_id = id

	stmt, err := db.Prepare("delete from sp_parcel_dls where parcel_dls_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_dls.Parcel_dls_id)
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


