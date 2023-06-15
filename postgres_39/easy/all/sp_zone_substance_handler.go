package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpZoneSubstance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_zone_substance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_zone_substance

	for rows.Next() {
		var sp_zone_substance dto.Sp_zone_substance
		if err := rows.Scan(&sp_zone_substance.Spatial_description_id, &sp_zone_substance.Spatial_obs_no, &sp_zone_substance.Mineral_zone_id, &sp_zone_substance.Substance_id, &sp_zone_substance.Active_ind, &sp_zone_substance.Base_depth, &sp_zone_substance.Base_depth_ouom, &sp_zone_substance.Base_qualifier, &sp_zone_substance.Base_strat_unit_id, &sp_zone_substance.Effective_date, &sp_zone_substance.Excluded_ind, &sp_zone_substance.Expiry_date, &sp_zone_substance.Included_ind, &sp_zone_substance.Ppdm_guid, &sp_zone_substance.Remark, &sp_zone_substance.Source, &sp_zone_substance.Strat_name_set_id, &sp_zone_substance.Top_depth, &sp_zone_substance.Top_depth_ouom, &sp_zone_substance.Top_qualifier, &sp_zone_substance.Top_strat_unit_id, &sp_zone_substance.Row_changed_by, &sp_zone_substance.Row_changed_date, &sp_zone_substance.Row_created_by, &sp_zone_substance.Row_created_date, &sp_zone_substance.Row_effective_date, &sp_zone_substance.Row_expiry_date, &sp_zone_substance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_zone_substance to result
		result = append(result, sp_zone_substance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpZoneSubstance(c *fiber.Ctx) error {
	var sp_zone_substance dto.Sp_zone_substance

	setDefaults(&sp_zone_substance)

	if err := json.Unmarshal(c.Body(), &sp_zone_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_zone_substance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	sp_zone_substance.Row_created_date = formatDate(sp_zone_substance.Row_created_date)
	sp_zone_substance.Row_changed_date = formatDate(sp_zone_substance.Row_changed_date)
	sp_zone_substance.Effective_date = formatDateString(sp_zone_substance.Effective_date)
	sp_zone_substance.Expiry_date = formatDateString(sp_zone_substance.Expiry_date)
	sp_zone_substance.Row_effective_date = formatDateString(sp_zone_substance.Row_effective_date)
	sp_zone_substance.Row_expiry_date = formatDateString(sp_zone_substance.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_substance.Spatial_description_id, sp_zone_substance.Spatial_obs_no, sp_zone_substance.Mineral_zone_id, sp_zone_substance.Substance_id, sp_zone_substance.Active_ind, sp_zone_substance.Base_depth, sp_zone_substance.Base_depth_ouom, sp_zone_substance.Base_qualifier, sp_zone_substance.Base_strat_unit_id, sp_zone_substance.Effective_date, sp_zone_substance.Excluded_ind, sp_zone_substance.Expiry_date, sp_zone_substance.Included_ind, sp_zone_substance.Ppdm_guid, sp_zone_substance.Remark, sp_zone_substance.Source, sp_zone_substance.Strat_name_set_id, sp_zone_substance.Top_depth, sp_zone_substance.Top_depth_ouom, sp_zone_substance.Top_qualifier, sp_zone_substance.Top_strat_unit_id, sp_zone_substance.Row_changed_by, sp_zone_substance.Row_changed_date, sp_zone_substance.Row_created_by, sp_zone_substance.Row_created_date, sp_zone_substance.Row_effective_date, sp_zone_substance.Row_expiry_date, sp_zone_substance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpZoneSubstance(c *fiber.Ctx) error {
	var sp_zone_substance dto.Sp_zone_substance

	setDefaults(&sp_zone_substance)

	if err := json.Unmarshal(c.Body(), &sp_zone_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_zone_substance.Spatial_description_id = id

    if sp_zone_substance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_zone_substance set spatial_obs_no = :1, mineral_zone_id = :2, substance_id = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, base_qualifier = :7, base_strat_unit_id = :8, effective_date = :9, excluded_ind = :10, expiry_date = :11, included_ind = :12, ppdm_guid = :13, remark = :14, source = :15, strat_name_set_id = :16, top_depth = :17, top_depth_ouom = :18, top_qualifier = :19, top_strat_unit_id = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where spatial_description_id = :28")
	if err != nil {
		return err
	}

	sp_zone_substance.Row_changed_date = formatDate(sp_zone_substance.Row_changed_date)
	sp_zone_substance.Effective_date = formatDateString(sp_zone_substance.Effective_date)
	sp_zone_substance.Expiry_date = formatDateString(sp_zone_substance.Expiry_date)
	sp_zone_substance.Row_effective_date = formatDateString(sp_zone_substance.Row_effective_date)
	sp_zone_substance.Row_expiry_date = formatDateString(sp_zone_substance.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_substance.Spatial_obs_no, sp_zone_substance.Mineral_zone_id, sp_zone_substance.Substance_id, sp_zone_substance.Active_ind, sp_zone_substance.Base_depth, sp_zone_substance.Base_depth_ouom, sp_zone_substance.Base_qualifier, sp_zone_substance.Base_strat_unit_id, sp_zone_substance.Effective_date, sp_zone_substance.Excluded_ind, sp_zone_substance.Expiry_date, sp_zone_substance.Included_ind, sp_zone_substance.Ppdm_guid, sp_zone_substance.Remark, sp_zone_substance.Source, sp_zone_substance.Strat_name_set_id, sp_zone_substance.Top_depth, sp_zone_substance.Top_depth_ouom, sp_zone_substance.Top_qualifier, sp_zone_substance.Top_strat_unit_id, sp_zone_substance.Row_changed_by, sp_zone_substance.Row_changed_date, sp_zone_substance.Row_created_by, sp_zone_substance.Row_effective_date, sp_zone_substance.Row_expiry_date, sp_zone_substance.Row_quality, sp_zone_substance.Spatial_description_id)
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

func PatchSpZoneSubstance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_zone_substance set "
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
	query += " where spatial_description_id = :id"

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

func DeleteSpZoneSubstance(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_zone_substance dto.Sp_zone_substance
	sp_zone_substance.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_zone_substance where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_zone_substance.Spatial_description_id)
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


