package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemGroup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_group")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_group

	for rows.Next() {
		var rm_info_item_group dto.Rm_info_item_group
		if err := rows.Scan(&rm_info_item_group.Group_info_item_subtype, &rm_info_item_group.Group_information_item_id, &rm_info_item_group.Info_item_subtype, &rm_info_item_group.Information_item_id, &rm_info_item_group.Group_obs_no, &rm_info_item_group.Active_ind, &rm_info_item_group.Effective_date, &rm_info_item_group.Expiry_date, &rm_info_item_group.Group_seq_no, &rm_info_item_group.Group_type, &rm_info_item_group.Ppdm_guid, &rm_info_item_group.Remark, &rm_info_item_group.Source, &rm_info_item_group.Row_changed_by, &rm_info_item_group.Row_changed_date, &rm_info_item_group.Row_created_by, &rm_info_item_group.Row_created_date, &rm_info_item_group.Row_effective_date, &rm_info_item_group.Row_expiry_date, &rm_info_item_group.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_group to result
		result = append(result, rm_info_item_group)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemGroup(c *fiber.Ctx) error {
	var rm_info_item_group dto.Rm_info_item_group

	setDefaults(&rm_info_item_group)

	if err := json.Unmarshal(c.Body(), &rm_info_item_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_group values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rm_info_item_group.Row_created_date = formatDate(rm_info_item_group.Row_created_date)
	rm_info_item_group.Row_changed_date = formatDate(rm_info_item_group.Row_changed_date)
	rm_info_item_group.Effective_date = formatDateString(rm_info_item_group.Effective_date)
	rm_info_item_group.Expiry_date = formatDateString(rm_info_item_group.Expiry_date)
	rm_info_item_group.Row_effective_date = formatDateString(rm_info_item_group.Row_effective_date)
	rm_info_item_group.Row_expiry_date = formatDateString(rm_info_item_group.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_group.Group_info_item_subtype, rm_info_item_group.Group_information_item_id, rm_info_item_group.Info_item_subtype, rm_info_item_group.Information_item_id, rm_info_item_group.Group_obs_no, rm_info_item_group.Active_ind, rm_info_item_group.Effective_date, rm_info_item_group.Expiry_date, rm_info_item_group.Group_seq_no, rm_info_item_group.Group_type, rm_info_item_group.Ppdm_guid, rm_info_item_group.Remark, rm_info_item_group.Source, rm_info_item_group.Row_changed_by, rm_info_item_group.Row_changed_date, rm_info_item_group.Row_created_by, rm_info_item_group.Row_created_date, rm_info_item_group.Row_effective_date, rm_info_item_group.Row_expiry_date, rm_info_item_group.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemGroup(c *fiber.Ctx) error {
	var rm_info_item_group dto.Rm_info_item_group

	setDefaults(&rm_info_item_group)

	if err := json.Unmarshal(c.Body(), &rm_info_item_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_group.Group_info_item_subtype = id

    if rm_info_item_group.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_group set group_information_item_id = :1, info_item_subtype = :2, information_item_id = :3, group_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, group_seq_no = :8, group_type = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where group_info_item_subtype = :20")
	if err != nil {
		return err
	}

	rm_info_item_group.Row_changed_date = formatDate(rm_info_item_group.Row_changed_date)
	rm_info_item_group.Effective_date = formatDateString(rm_info_item_group.Effective_date)
	rm_info_item_group.Expiry_date = formatDateString(rm_info_item_group.Expiry_date)
	rm_info_item_group.Row_effective_date = formatDateString(rm_info_item_group.Row_effective_date)
	rm_info_item_group.Row_expiry_date = formatDateString(rm_info_item_group.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_group.Group_information_item_id, rm_info_item_group.Info_item_subtype, rm_info_item_group.Information_item_id, rm_info_item_group.Group_obs_no, rm_info_item_group.Active_ind, rm_info_item_group.Effective_date, rm_info_item_group.Expiry_date, rm_info_item_group.Group_seq_no, rm_info_item_group.Group_type, rm_info_item_group.Ppdm_guid, rm_info_item_group.Remark, rm_info_item_group.Source, rm_info_item_group.Row_changed_by, rm_info_item_group.Row_changed_date, rm_info_item_group.Row_created_by, rm_info_item_group.Row_effective_date, rm_info_item_group.Row_expiry_date, rm_info_item_group.Row_quality, rm_info_item_group.Group_info_item_subtype)
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

func PatchRmInfoItemGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_group set "
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
	query += " where group_info_item_subtype = :id"

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

func DeleteRmInfoItemGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_group dto.Rm_info_item_group
	rm_info_item_group.Group_info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_group where group_info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_group.Group_info_item_subtype)
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


