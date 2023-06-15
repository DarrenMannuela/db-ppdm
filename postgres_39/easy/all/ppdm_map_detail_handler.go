package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMapDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_map_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_map_detail

	for rows.Next() {
		var ppdm_map_detail dto.Ppdm_map_detail
		if err := rows.Scan(&ppdm_map_detail.Map_id, &ppdm_map_detail.Map_detail_obs_no, &ppdm_map_detail.Active_ind, &ppdm_map_detail.Column_name1, &ppdm_map_detail.Column_name2, &ppdm_map_detail.Create_method, &ppdm_map_detail.Default_value, &ppdm_map_detail.Effective_date, &ppdm_map_detail.Expiry_date, &ppdm_map_detail.Map_desc, &ppdm_map_detail.Map_owner_ba_id, &ppdm_map_detail.Map_type, &ppdm_map_detail.Map_version_num, &ppdm_map_detail.Ppdm_guid, &ppdm_map_detail.Preferred_ind, &ppdm_map_detail.Remark, &ppdm_map_detail.Ring_seq_no, &ppdm_map_detail.Schema_entity_id1, &ppdm_map_detail.Schema_entity_id2, &ppdm_map_detail.Source, &ppdm_map_detail.Sw_application_id, &ppdm_map_detail.System_id1, &ppdm_map_detail.System_id2, &ppdm_map_detail.Table_name1, &ppdm_map_detail.Table_name2, &ppdm_map_detail.Row_changed_by, &ppdm_map_detail.Row_changed_date, &ppdm_map_detail.Row_created_by, &ppdm_map_detail.Row_created_date, &ppdm_map_detail.Row_effective_date, &ppdm_map_detail.Row_expiry_date, &ppdm_map_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_map_detail to result
		result = append(result, ppdm_map_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMapDetail(c *fiber.Ctx) error {
	var ppdm_map_detail dto.Ppdm_map_detail

	setDefaults(&ppdm_map_detail)

	if err := json.Unmarshal(c.Body(), &ppdm_map_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_map_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ppdm_map_detail.Row_created_date = formatDate(ppdm_map_detail.Row_created_date)
	ppdm_map_detail.Row_changed_date = formatDate(ppdm_map_detail.Row_changed_date)
	ppdm_map_detail.Effective_date = formatDateString(ppdm_map_detail.Effective_date)
	ppdm_map_detail.Expiry_date = formatDateString(ppdm_map_detail.Expiry_date)
	ppdm_map_detail.Row_effective_date = formatDateString(ppdm_map_detail.Row_effective_date)
	ppdm_map_detail.Row_expiry_date = formatDateString(ppdm_map_detail.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_detail.Map_id, ppdm_map_detail.Map_detail_obs_no, ppdm_map_detail.Active_ind, ppdm_map_detail.Column_name1, ppdm_map_detail.Column_name2, ppdm_map_detail.Create_method, ppdm_map_detail.Default_value, ppdm_map_detail.Effective_date, ppdm_map_detail.Expiry_date, ppdm_map_detail.Map_desc, ppdm_map_detail.Map_owner_ba_id, ppdm_map_detail.Map_type, ppdm_map_detail.Map_version_num, ppdm_map_detail.Ppdm_guid, ppdm_map_detail.Preferred_ind, ppdm_map_detail.Remark, ppdm_map_detail.Ring_seq_no, ppdm_map_detail.Schema_entity_id1, ppdm_map_detail.Schema_entity_id2, ppdm_map_detail.Source, ppdm_map_detail.Sw_application_id, ppdm_map_detail.System_id1, ppdm_map_detail.System_id2, ppdm_map_detail.Table_name1, ppdm_map_detail.Table_name2, ppdm_map_detail.Row_changed_by, ppdm_map_detail.Row_changed_date, ppdm_map_detail.Row_created_by, ppdm_map_detail.Row_created_date, ppdm_map_detail.Row_effective_date, ppdm_map_detail.Row_expiry_date, ppdm_map_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMapDetail(c *fiber.Ctx) error {
	var ppdm_map_detail dto.Ppdm_map_detail

	setDefaults(&ppdm_map_detail)

	if err := json.Unmarshal(c.Body(), &ppdm_map_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_map_detail.Map_id = id

    if ppdm_map_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_map_detail set map_detail_obs_no = :1, active_ind = :2, column_name1 = :3, column_name2 = :4, create_method = :5, default_value = :6, effective_date = :7, expiry_date = :8, map_desc = :9, map_owner_ba_id = :10, map_type = :11, map_version_num = :12, ppdm_guid = :13, preferred_ind = :14, remark = :15, ring_seq_no = :16, schema_entity_id1 = :17, schema_entity_id2 = :18, source = :19, sw_application_id = :20, system_id1 = :21, system_id2 = :22, table_name1 = :23, table_name2 = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where map_id = :32")
	if err != nil {
		return err
	}

	ppdm_map_detail.Row_changed_date = formatDate(ppdm_map_detail.Row_changed_date)
	ppdm_map_detail.Effective_date = formatDateString(ppdm_map_detail.Effective_date)
	ppdm_map_detail.Expiry_date = formatDateString(ppdm_map_detail.Expiry_date)
	ppdm_map_detail.Row_effective_date = formatDateString(ppdm_map_detail.Row_effective_date)
	ppdm_map_detail.Row_expiry_date = formatDateString(ppdm_map_detail.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_detail.Map_detail_obs_no, ppdm_map_detail.Active_ind, ppdm_map_detail.Column_name1, ppdm_map_detail.Column_name2, ppdm_map_detail.Create_method, ppdm_map_detail.Default_value, ppdm_map_detail.Effective_date, ppdm_map_detail.Expiry_date, ppdm_map_detail.Map_desc, ppdm_map_detail.Map_owner_ba_id, ppdm_map_detail.Map_type, ppdm_map_detail.Map_version_num, ppdm_map_detail.Ppdm_guid, ppdm_map_detail.Preferred_ind, ppdm_map_detail.Remark, ppdm_map_detail.Ring_seq_no, ppdm_map_detail.Schema_entity_id1, ppdm_map_detail.Schema_entity_id2, ppdm_map_detail.Source, ppdm_map_detail.Sw_application_id, ppdm_map_detail.System_id1, ppdm_map_detail.System_id2, ppdm_map_detail.Table_name1, ppdm_map_detail.Table_name2, ppdm_map_detail.Row_changed_by, ppdm_map_detail.Row_changed_date, ppdm_map_detail.Row_created_by, ppdm_map_detail.Row_effective_date, ppdm_map_detail.Row_expiry_date, ppdm_map_detail.Row_quality, ppdm_map_detail.Map_id)
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

func PatchPpdmMapDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_map_detail set "
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
	query += " where map_id = :id"

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

func DeletePpdmMapDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_map_detail dto.Ppdm_map_detail
	ppdm_map_detail.Map_id = id

	stmt, err := db.Prepare("delete from ppdm_map_detail where map_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_map_detail.Map_id)
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


