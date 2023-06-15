package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratTopoRelation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_topo_relation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_topo_relation

	for rows.Next() {
		var strat_topo_relation dto.Strat_topo_relation
		if err := rows.Scan(&strat_topo_relation.Strat_name_set_id_1, &strat_topo_relation.Strat_unit_id_1, &strat_topo_relation.Strat_name_set_id_2, &strat_topo_relation.Strat_unit_id_2, &strat_topo_relation.Source, &strat_topo_relation.Topo_relation_id, &strat_topo_relation.Active_ind, &strat_topo_relation.Area_id, &strat_topo_relation.Area_type, &strat_topo_relation.Effective_date, &strat_topo_relation.Expiry_date, &strat_topo_relation.Interp_id_1, &strat_topo_relation.Interp_id_2, &strat_topo_relation.Ppdm_guid, &strat_topo_relation.Remark, &strat_topo_relation.Strat_column_id_1, &strat_topo_relation.Strat_column_id_2, &strat_topo_relation.Strat_column_source_1, &strat_topo_relation.Strat_column_source_2, &strat_topo_relation.Topo_relation_type, &strat_topo_relation.Row_changed_by, &strat_topo_relation.Row_changed_date, &strat_topo_relation.Row_created_by, &strat_topo_relation.Row_created_date, &strat_topo_relation.Row_effective_date, &strat_topo_relation.Row_expiry_date, &strat_topo_relation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_topo_relation to result
		result = append(result, strat_topo_relation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratTopoRelation(c *fiber.Ctx) error {
	var strat_topo_relation dto.Strat_topo_relation

	setDefaults(&strat_topo_relation)

	if err := json.Unmarshal(c.Body(), &strat_topo_relation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_topo_relation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	strat_topo_relation.Row_created_date = formatDate(strat_topo_relation.Row_created_date)
	strat_topo_relation.Row_changed_date = formatDate(strat_topo_relation.Row_changed_date)
	strat_topo_relation.Effective_date = formatDateString(strat_topo_relation.Effective_date)
	strat_topo_relation.Expiry_date = formatDateString(strat_topo_relation.Expiry_date)
	strat_topo_relation.Row_effective_date = formatDateString(strat_topo_relation.Row_effective_date)
	strat_topo_relation.Row_expiry_date = formatDateString(strat_topo_relation.Row_expiry_date)






	rows, err := stmt.Exec(strat_topo_relation.Strat_name_set_id_1, strat_topo_relation.Strat_unit_id_1, strat_topo_relation.Strat_name_set_id_2, strat_topo_relation.Strat_unit_id_2, strat_topo_relation.Source, strat_topo_relation.Topo_relation_id, strat_topo_relation.Active_ind, strat_topo_relation.Area_id, strat_topo_relation.Area_type, strat_topo_relation.Effective_date, strat_topo_relation.Expiry_date, strat_topo_relation.Interp_id_1, strat_topo_relation.Interp_id_2, strat_topo_relation.Ppdm_guid, strat_topo_relation.Remark, strat_topo_relation.Strat_column_id_1, strat_topo_relation.Strat_column_id_2, strat_topo_relation.Strat_column_source_1, strat_topo_relation.Strat_column_source_2, strat_topo_relation.Topo_relation_type, strat_topo_relation.Row_changed_by, strat_topo_relation.Row_changed_date, strat_topo_relation.Row_created_by, strat_topo_relation.Row_created_date, strat_topo_relation.Row_effective_date, strat_topo_relation.Row_expiry_date, strat_topo_relation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratTopoRelation(c *fiber.Ctx) error {
	var strat_topo_relation dto.Strat_topo_relation

	setDefaults(&strat_topo_relation)

	if err := json.Unmarshal(c.Body(), &strat_topo_relation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_topo_relation.Strat_name_set_id_1 = id

    if strat_topo_relation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_topo_relation set strat_unit_id_1 = :1, strat_name_set_id_2 = :2, strat_unit_id_2 = :3, source = :4, topo_relation_id = :5, active_ind = :6, area_id = :7, area_type = :8, effective_date = :9, expiry_date = :10, interp_id_1 = :11, interp_id_2 = :12, ppdm_guid = :13, remark = :14, strat_column_id_1 = :15, strat_column_id_2 = :16, strat_column_source_1 = :17, strat_column_source_2 = :18, topo_relation_type = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where strat_name_set_id_1 = :27")
	if err != nil {
		return err
	}

	strat_topo_relation.Row_changed_date = formatDate(strat_topo_relation.Row_changed_date)
	strat_topo_relation.Effective_date = formatDateString(strat_topo_relation.Effective_date)
	strat_topo_relation.Expiry_date = formatDateString(strat_topo_relation.Expiry_date)
	strat_topo_relation.Row_effective_date = formatDateString(strat_topo_relation.Row_effective_date)
	strat_topo_relation.Row_expiry_date = formatDateString(strat_topo_relation.Row_expiry_date)






	rows, err := stmt.Exec(strat_topo_relation.Strat_unit_id_1, strat_topo_relation.Strat_name_set_id_2, strat_topo_relation.Strat_unit_id_2, strat_topo_relation.Source, strat_topo_relation.Topo_relation_id, strat_topo_relation.Active_ind, strat_topo_relation.Area_id, strat_topo_relation.Area_type, strat_topo_relation.Effective_date, strat_topo_relation.Expiry_date, strat_topo_relation.Interp_id_1, strat_topo_relation.Interp_id_2, strat_topo_relation.Ppdm_guid, strat_topo_relation.Remark, strat_topo_relation.Strat_column_id_1, strat_topo_relation.Strat_column_id_2, strat_topo_relation.Strat_column_source_1, strat_topo_relation.Strat_column_source_2, strat_topo_relation.Topo_relation_type, strat_topo_relation.Row_changed_by, strat_topo_relation.Row_changed_date, strat_topo_relation.Row_created_by, strat_topo_relation.Row_effective_date, strat_topo_relation.Row_expiry_date, strat_topo_relation.Row_quality, strat_topo_relation.Strat_name_set_id_1)
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

func PatchStratTopoRelation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_topo_relation set "
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
	query += " where strat_name_set_id_1 = :id"

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

func DeleteStratTopoRelation(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_topo_relation dto.Strat_topo_relation
	strat_topo_relation.Strat_name_set_id_1 = id

	stmt, err := db.Prepare("delete from strat_topo_relation where strat_name_set_id_1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_topo_relation.Strat_name_set_id_1)
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


