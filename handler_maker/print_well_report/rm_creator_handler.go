package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmCreator(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_creator")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_creator

	for rows.Next() {
		var rm_creator dto.Rm_creator
		if err := rows.Scan(&rm_creator.Info_item_subtype, &rm_creator.Information_item_id, &rm_creator.Creator_id, &rm_creator.Active_ind, &rm_creator.Creator_ba_id, &rm_creator.Creator_ba_type, &rm_creator.Creator_name, &rm_creator.Creator_type, &rm_creator.Effective_date, &rm_creator.Expiry_date, &rm_creator.First_name, &rm_creator.Last_name, &rm_creator.Middle_name, &rm_creator.Ppdm_guid, &rm_creator.Remark, &rm_creator.Source, &rm_creator.Row_changed_by, &rm_creator.Row_changed_date, &rm_creator.Row_created_by, &rm_creator.Row_created_date, &rm_creator.Row_effective_date, &rm_creator.Row_expiry_date, &rm_creator.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_creator to result
		result = append(result, rm_creator)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmCreator(c *fiber.Ctx) error {
	var rm_creator dto.Rm_creator

	setDefaults(&rm_creator)

	if err := json.Unmarshal(c.Body(), &rm_creator); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_creator values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	rm_creator.Row_created_date = formatDate(rm_creator.Row_created_date)
	rm_creator.Row_changed_date = formatDate(rm_creator.Row_changed_date)
	rm_creator.Effective_date = formatDateString(rm_creator.Effective_date)
	rm_creator.Expiry_date = formatDateString(rm_creator.Expiry_date)
	rm_creator.Row_effective_date = formatDateString(rm_creator.Row_effective_date)
	rm_creator.Row_expiry_date = formatDateString(rm_creator.Row_expiry_date)






	rows, err := stmt.Exec(rm_creator.Info_item_subtype, rm_creator.Information_item_id, rm_creator.Creator_id, rm_creator.Active_ind, rm_creator.Creator_ba_id, rm_creator.Creator_ba_type, rm_creator.Creator_name, rm_creator.Creator_type, rm_creator.Effective_date, rm_creator.Expiry_date, rm_creator.First_name, rm_creator.Last_name, rm_creator.Middle_name, rm_creator.Ppdm_guid, rm_creator.Remark, rm_creator.Source, rm_creator.Row_changed_by, rm_creator.Row_changed_date, rm_creator.Row_created_by, rm_creator.Row_created_date, rm_creator.Row_effective_date, rm_creator.Row_expiry_date, rm_creator.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmCreator(c *fiber.Ctx) error {
	var rm_creator dto.Rm_creator

	setDefaults(&rm_creator)

	if err := json.Unmarshal(c.Body(), &rm_creator); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_creator.Info_item_subtype = id

    if rm_creator.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_creator set information_item_id = :1, creator_id = :2, active_ind = :3, creator_ba_id = :4, creator_ba_type = :5, creator_name = :6, creator_type = :7, effective_date = :8, expiry_date = :9, first_name = :10, last_name = :11, middle_name = :12, ppdm_guid = :13, remark = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where info_item_subtype = :23")
	if err != nil {
		return err
	}

	rm_creator.Row_changed_date = formatDate(rm_creator.Row_changed_date)
	rm_creator.Effective_date = formatDateString(rm_creator.Effective_date)
	rm_creator.Expiry_date = formatDateString(rm_creator.Expiry_date)
	rm_creator.Row_effective_date = formatDateString(rm_creator.Row_effective_date)
	rm_creator.Row_expiry_date = formatDateString(rm_creator.Row_expiry_date)






	rows, err := stmt.Exec(rm_creator.Information_item_id, rm_creator.Creator_id, rm_creator.Active_ind, rm_creator.Creator_ba_id, rm_creator.Creator_ba_type, rm_creator.Creator_name, rm_creator.Creator_type, rm_creator.Effective_date, rm_creator.Expiry_date, rm_creator.First_name, rm_creator.Last_name, rm_creator.Middle_name, rm_creator.Ppdm_guid, rm_creator.Remark, rm_creator.Source, rm_creator.Row_changed_by, rm_creator.Row_changed_date, rm_creator.Row_created_by, rm_creator.Row_effective_date, rm_creator.Row_expiry_date, rm_creator.Row_quality, rm_creator.Info_item_subtype)
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

func PatchRmCreator(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_creator set "
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
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where info_item_subtype = :id"

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

func DeleteRmCreator(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_creator dto.Rm_creator
	rm_creator.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_creator where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_creator.Info_item_subtype)
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


