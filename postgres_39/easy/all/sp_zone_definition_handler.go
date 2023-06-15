package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpZoneDefinition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_zone_definition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_zone_definition

	for rows.Next() {
		var sp_zone_definition dto.Sp_zone_definition
		if err := rows.Scan(&sp_zone_definition.Zone_definition_id, &sp_zone_definition.Active_ind, &sp_zone_definition.Base_depth, &sp_zone_definition.Base_depth_ouom, &sp_zone_definition.Base_qualifier, &sp_zone_definition.Base_strat_unit_id, &sp_zone_definition.Defined_strat_unit_id, &sp_zone_definition.Description, &sp_zone_definition.Effective_date, &sp_zone_definition.Expiry_date, &sp_zone_definition.Owner_ba_id, &sp_zone_definition.Owner_ref_num, &sp_zone_definition.Ppdm_guid, &sp_zone_definition.Remark, &sp_zone_definition.Source, &sp_zone_definition.Strat_name_set_id, &sp_zone_definition.Top_depth, &sp_zone_definition.Top_depth_ouom, &sp_zone_definition.Top_qualifier, &sp_zone_definition.Top_strat_unit_id, &sp_zone_definition.Uwi, &sp_zone_definition.Zone_derivation_method, &sp_zone_definition.Zone_type, &sp_zone_definition.Row_changed_by, &sp_zone_definition.Row_changed_date, &sp_zone_definition.Row_created_by, &sp_zone_definition.Row_created_date, &sp_zone_definition.Row_effective_date, &sp_zone_definition.Row_expiry_date, &sp_zone_definition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_zone_definition to result
		result = append(result, sp_zone_definition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpZoneDefinition(c *fiber.Ctx) error {
	var sp_zone_definition dto.Sp_zone_definition

	setDefaults(&sp_zone_definition)

	if err := json.Unmarshal(c.Body(), &sp_zone_definition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_zone_definition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	sp_zone_definition.Row_created_date = formatDate(sp_zone_definition.Row_created_date)
	sp_zone_definition.Row_changed_date = formatDate(sp_zone_definition.Row_changed_date)
	sp_zone_definition.Effective_date = formatDateString(sp_zone_definition.Effective_date)
	sp_zone_definition.Expiry_date = formatDateString(sp_zone_definition.Expiry_date)
	sp_zone_definition.Row_effective_date = formatDateString(sp_zone_definition.Row_effective_date)
	sp_zone_definition.Row_expiry_date = formatDateString(sp_zone_definition.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_definition.Zone_definition_id, sp_zone_definition.Active_ind, sp_zone_definition.Base_depth, sp_zone_definition.Base_depth_ouom, sp_zone_definition.Base_qualifier, sp_zone_definition.Base_strat_unit_id, sp_zone_definition.Defined_strat_unit_id, sp_zone_definition.Description, sp_zone_definition.Effective_date, sp_zone_definition.Expiry_date, sp_zone_definition.Owner_ba_id, sp_zone_definition.Owner_ref_num, sp_zone_definition.Ppdm_guid, sp_zone_definition.Remark, sp_zone_definition.Source, sp_zone_definition.Strat_name_set_id, sp_zone_definition.Top_depth, sp_zone_definition.Top_depth_ouom, sp_zone_definition.Top_qualifier, sp_zone_definition.Top_strat_unit_id, sp_zone_definition.Uwi, sp_zone_definition.Zone_derivation_method, sp_zone_definition.Zone_type, sp_zone_definition.Row_changed_by, sp_zone_definition.Row_changed_date, sp_zone_definition.Row_created_by, sp_zone_definition.Row_created_date, sp_zone_definition.Row_effective_date, sp_zone_definition.Row_expiry_date, sp_zone_definition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpZoneDefinition(c *fiber.Ctx) error {
	var sp_zone_definition dto.Sp_zone_definition

	setDefaults(&sp_zone_definition)

	if err := json.Unmarshal(c.Body(), &sp_zone_definition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_zone_definition.Zone_definition_id = id

    if sp_zone_definition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_zone_definition set active_ind = :1, base_depth = :2, base_depth_ouom = :3, base_qualifier = :4, base_strat_unit_id = :5, defined_strat_unit_id = :6, description = :7, effective_date = :8, expiry_date = :9, owner_ba_id = :10, owner_ref_num = :11, ppdm_guid = :12, remark = :13, source = :14, strat_name_set_id = :15, top_depth = :16, top_depth_ouom = :17, top_qualifier = :18, top_strat_unit_id = :19, uwi = :20, zone_derivation_method = :21, zone_type = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where zone_definition_id = :30")
	if err != nil {
		return err
	}

	sp_zone_definition.Row_changed_date = formatDate(sp_zone_definition.Row_changed_date)
	sp_zone_definition.Effective_date = formatDateString(sp_zone_definition.Effective_date)
	sp_zone_definition.Expiry_date = formatDateString(sp_zone_definition.Expiry_date)
	sp_zone_definition.Row_effective_date = formatDateString(sp_zone_definition.Row_effective_date)
	sp_zone_definition.Row_expiry_date = formatDateString(sp_zone_definition.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_definition.Active_ind, sp_zone_definition.Base_depth, sp_zone_definition.Base_depth_ouom, sp_zone_definition.Base_qualifier, sp_zone_definition.Base_strat_unit_id, sp_zone_definition.Defined_strat_unit_id, sp_zone_definition.Description, sp_zone_definition.Effective_date, sp_zone_definition.Expiry_date, sp_zone_definition.Owner_ba_id, sp_zone_definition.Owner_ref_num, sp_zone_definition.Ppdm_guid, sp_zone_definition.Remark, sp_zone_definition.Source, sp_zone_definition.Strat_name_set_id, sp_zone_definition.Top_depth, sp_zone_definition.Top_depth_ouom, sp_zone_definition.Top_qualifier, sp_zone_definition.Top_strat_unit_id, sp_zone_definition.Uwi, sp_zone_definition.Zone_derivation_method, sp_zone_definition.Zone_type, sp_zone_definition.Row_changed_by, sp_zone_definition.Row_changed_date, sp_zone_definition.Row_created_by, sp_zone_definition.Row_effective_date, sp_zone_definition.Row_expiry_date, sp_zone_definition.Row_quality, sp_zone_definition.Zone_definition_id)
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

func PatchSpZoneDefinition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_zone_definition set "
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
	query += " where zone_definition_id = :id"

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

func DeleteSpZoneDefinition(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_zone_definition dto.Sp_zone_definition
	sp_zone_definition.Zone_definition_id = id

	stmt, err := db.Prepare("delete from sp_zone_definition where zone_definition_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_zone_definition.Zone_definition_id)
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


