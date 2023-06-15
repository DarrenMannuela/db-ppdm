package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpMineralZone(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_mineral_zone")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_mineral_zone

	for rows.Next() {
		var sp_mineral_zone dto.Sp_mineral_zone
		if err := rows.Scan(&sp_mineral_zone.Spatial_description_id, &sp_mineral_zone.Spatial_obs_no, &sp_mineral_zone.Mineral_zone_id, &sp_mineral_zone.Active_ind, &sp_mineral_zone.Base_qualifier, &sp_mineral_zone.Base_zone_definition_id, &sp_mineral_zone.Deep_right_reversion_ind, &sp_mineral_zone.Effective_date, &sp_mineral_zone.Expiry_date, &sp_mineral_zone.Inactivation_date, &sp_mineral_zone.Metes_and_bounds, &sp_mineral_zone.Ppdm_guid, &sp_mineral_zone.Remark, &sp_mineral_zone.Source, &sp_mineral_zone.Top_qualifier, &sp_mineral_zone.Top_zone_definition_id, &sp_mineral_zone.Row_changed_by, &sp_mineral_zone.Row_changed_date, &sp_mineral_zone.Row_created_by, &sp_mineral_zone.Row_created_date, &sp_mineral_zone.Row_effective_date, &sp_mineral_zone.Row_expiry_date, &sp_mineral_zone.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_mineral_zone to result
		result = append(result, sp_mineral_zone)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpMineralZone(c *fiber.Ctx) error {
	var sp_mineral_zone dto.Sp_mineral_zone

	setDefaults(&sp_mineral_zone)

	if err := json.Unmarshal(c.Body(), &sp_mineral_zone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_mineral_zone values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	sp_mineral_zone.Row_created_date = formatDate(sp_mineral_zone.Row_created_date)
	sp_mineral_zone.Row_changed_date = formatDate(sp_mineral_zone.Row_changed_date)
	sp_mineral_zone.Effective_date = formatDateString(sp_mineral_zone.Effective_date)
	sp_mineral_zone.Expiry_date = formatDateString(sp_mineral_zone.Expiry_date)
	sp_mineral_zone.Inactivation_date = formatDateString(sp_mineral_zone.Inactivation_date)
	sp_mineral_zone.Row_effective_date = formatDateString(sp_mineral_zone.Row_effective_date)
	sp_mineral_zone.Row_expiry_date = formatDateString(sp_mineral_zone.Row_expiry_date)






	rows, err := stmt.Exec(sp_mineral_zone.Spatial_description_id, sp_mineral_zone.Spatial_obs_no, sp_mineral_zone.Mineral_zone_id, sp_mineral_zone.Active_ind, sp_mineral_zone.Base_qualifier, sp_mineral_zone.Base_zone_definition_id, sp_mineral_zone.Deep_right_reversion_ind, sp_mineral_zone.Effective_date, sp_mineral_zone.Expiry_date, sp_mineral_zone.Inactivation_date, sp_mineral_zone.Metes_and_bounds, sp_mineral_zone.Ppdm_guid, sp_mineral_zone.Remark, sp_mineral_zone.Source, sp_mineral_zone.Top_qualifier, sp_mineral_zone.Top_zone_definition_id, sp_mineral_zone.Row_changed_by, sp_mineral_zone.Row_changed_date, sp_mineral_zone.Row_created_by, sp_mineral_zone.Row_created_date, sp_mineral_zone.Row_effective_date, sp_mineral_zone.Row_expiry_date, sp_mineral_zone.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpMineralZone(c *fiber.Ctx) error {
	var sp_mineral_zone dto.Sp_mineral_zone

	setDefaults(&sp_mineral_zone)

	if err := json.Unmarshal(c.Body(), &sp_mineral_zone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_mineral_zone.Spatial_description_id = id

    if sp_mineral_zone.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_mineral_zone set spatial_obs_no = :1, mineral_zone_id = :2, active_ind = :3, base_qualifier = :4, base_zone_definition_id = :5, deep_right_reversion_ind = :6, effective_date = :7, expiry_date = :8, inactivation_date = :9, metes_and_bounds = :10, ppdm_guid = :11, remark = :12, source = :13, top_qualifier = :14, top_zone_definition_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where spatial_description_id = :23")
	if err != nil {
		return err
	}

	sp_mineral_zone.Row_changed_date = formatDate(sp_mineral_zone.Row_changed_date)
	sp_mineral_zone.Effective_date = formatDateString(sp_mineral_zone.Effective_date)
	sp_mineral_zone.Expiry_date = formatDateString(sp_mineral_zone.Expiry_date)
	sp_mineral_zone.Inactivation_date = formatDateString(sp_mineral_zone.Inactivation_date)
	sp_mineral_zone.Row_effective_date = formatDateString(sp_mineral_zone.Row_effective_date)
	sp_mineral_zone.Row_expiry_date = formatDateString(sp_mineral_zone.Row_expiry_date)






	rows, err := stmt.Exec(sp_mineral_zone.Spatial_obs_no, sp_mineral_zone.Mineral_zone_id, sp_mineral_zone.Active_ind, sp_mineral_zone.Base_qualifier, sp_mineral_zone.Base_zone_definition_id, sp_mineral_zone.Deep_right_reversion_ind, sp_mineral_zone.Effective_date, sp_mineral_zone.Expiry_date, sp_mineral_zone.Inactivation_date, sp_mineral_zone.Metes_and_bounds, sp_mineral_zone.Ppdm_guid, sp_mineral_zone.Remark, sp_mineral_zone.Source, sp_mineral_zone.Top_qualifier, sp_mineral_zone.Top_zone_definition_id, sp_mineral_zone.Row_changed_by, sp_mineral_zone.Row_changed_date, sp_mineral_zone.Row_created_by, sp_mineral_zone.Row_effective_date, sp_mineral_zone.Row_expiry_date, sp_mineral_zone.Row_quality, sp_mineral_zone.Spatial_description_id)
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

func PatchSpMineralZone(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_mineral_zone set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactivation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSpMineralZone(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_mineral_zone dto.Sp_mineral_zone
	sp_mineral_zone.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_mineral_zone where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_mineral_zone.Spatial_description_id)
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


