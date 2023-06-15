package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInformationItem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_information_item")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_information_item

	for rows.Next() {
		var rm_information_item dto.Rm_information_item
		if err := rows.Scan(&rm_information_item.Info_item_subtype, &rm_information_item.Information_item_id, &rm_information_item.Abstract, &rm_information_item.Accepted_date, &rm_information_item.Access_condition, &rm_information_item.Active_ind, &rm_information_item.Available_date, &rm_information_item.Coord_acquisition_id, &rm_information_item.Copyright_date, &rm_information_item.Effective_date, &rm_information_item.Expiry_date, &rm_information_item.Geog_coord_system_id, &rm_information_item.Group_ind, &rm_information_item.Issue_date, &rm_information_item.Item_category, &rm_information_item.Item_sub_category, &rm_information_item.Local_coord_system_id, &rm_information_item.Map_coord_system_id, &rm_information_item.Max_latitude, &rm_information_item.Max_longitude, &rm_information_item.Min_latitude, &rm_information_item.Min_longitude, &rm_information_item.Modified_date, &rm_information_item.Origin_date, &rm_information_item.Ppdm_guid, &rm_information_item.Publish_date, &rm_information_item.Purpose, &rm_information_item.Reference_num, &rm_information_item.Remark, &rm_information_item.Security_desc, &rm_information_item.Source, &rm_information_item.Source_document_id, &rm_information_item.Submit_date, &rm_information_item.Time_period_desc, &rm_information_item.Title, &rm_information_item.Use_condition, &rm_information_item.Version_num, &rm_information_item.Row_changed_by, &rm_information_item.Row_changed_date, &rm_information_item.Row_created_by, &rm_information_item.Row_created_date, &rm_information_item.Row_effective_date, &rm_information_item.Row_expiry_date, &rm_information_item.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_information_item to result
		result = append(result, rm_information_item)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInformationItem(c *fiber.Ctx) error {
	var rm_information_item dto.Rm_information_item

	setDefaults(&rm_information_item)

	if err := json.Unmarshal(c.Body(), &rm_information_item); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_information_item values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	rm_information_item.Row_created_date = formatDate(rm_information_item.Row_created_date)
	rm_information_item.Row_changed_date = formatDate(rm_information_item.Row_changed_date)
	rm_information_item.Accepted_date = formatDateString(rm_information_item.Accepted_date)
	rm_information_item.Available_date = formatDateString(rm_information_item.Available_date)
	rm_information_item.Copyright_date = formatDateString(rm_information_item.Copyright_date)
	rm_information_item.Effective_date = formatDateString(rm_information_item.Effective_date)
	rm_information_item.Expiry_date = formatDateString(rm_information_item.Expiry_date)
	rm_information_item.Issue_date = formatDateString(rm_information_item.Issue_date)
	rm_information_item.Modified_date = formatDateString(rm_information_item.Modified_date)
	rm_information_item.Origin_date = formatDateString(rm_information_item.Origin_date)
	rm_information_item.Publish_date = formatDateString(rm_information_item.Publish_date)
	rm_information_item.Submit_date = formatDateString(rm_information_item.Submit_date)
	rm_information_item.Row_effective_date = formatDateString(rm_information_item.Row_effective_date)
	rm_information_item.Row_expiry_date = formatDateString(rm_information_item.Row_expiry_date)






	rows, err := stmt.Exec(rm_information_item.Info_item_subtype, rm_information_item.Information_item_id, rm_information_item.Abstract, rm_information_item.Accepted_date, rm_information_item.Access_condition, rm_information_item.Active_ind, rm_information_item.Available_date, rm_information_item.Coord_acquisition_id, rm_information_item.Copyright_date, rm_information_item.Effective_date, rm_information_item.Expiry_date, rm_information_item.Geog_coord_system_id, rm_information_item.Group_ind, rm_information_item.Issue_date, rm_information_item.Item_category, rm_information_item.Item_sub_category, rm_information_item.Local_coord_system_id, rm_information_item.Map_coord_system_id, rm_information_item.Max_latitude, rm_information_item.Max_longitude, rm_information_item.Min_latitude, rm_information_item.Min_longitude, rm_information_item.Modified_date, rm_information_item.Origin_date, rm_information_item.Ppdm_guid, rm_information_item.Publish_date, rm_information_item.Purpose, rm_information_item.Reference_num, rm_information_item.Remark, rm_information_item.Security_desc, rm_information_item.Source, rm_information_item.Source_document_id, rm_information_item.Submit_date, rm_information_item.Time_period_desc, rm_information_item.Title, rm_information_item.Use_condition, rm_information_item.Version_num, rm_information_item.Row_changed_by, rm_information_item.Row_changed_date, rm_information_item.Row_created_by, rm_information_item.Row_created_date, rm_information_item.Row_effective_date, rm_information_item.Row_expiry_date, rm_information_item.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInformationItem(c *fiber.Ctx) error {
	var rm_information_item dto.Rm_information_item

	setDefaults(&rm_information_item)

	if err := json.Unmarshal(c.Body(), &rm_information_item); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_information_item.Info_item_subtype = id

    if rm_information_item.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_information_item set information_item_id = :1, abstract = :2, accepted_date = :3, access_condition = :4, active_ind = :5, available_date = :6, coord_acquisition_id = :7, copyright_date = :8, effective_date = :9, expiry_date = :10, geog_coord_system_id = :11, group_ind = :12, issue_date = :13, item_category = :14, item_sub_category = :15, local_coord_system_id = :16, map_coord_system_id = :17, max_latitude = :18, max_longitude = :19, min_latitude = :20, min_longitude = :21, modified_date = :22, origin_date = :23, ppdm_guid = :24, publish_date = :25, purpose = :26, reference_num = :27, remark = :28, security_desc = :29, source = :30, source_document_id = :31, submit_date = :32, time_period_desc = :33, title = :34, use_condition = :35, version_num = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where info_item_subtype = :44")
	if err != nil {
		return err
	}

	rm_information_item.Row_changed_date = formatDate(rm_information_item.Row_changed_date)
	rm_information_item.Accepted_date = formatDateString(rm_information_item.Accepted_date)
	rm_information_item.Available_date = formatDateString(rm_information_item.Available_date)
	rm_information_item.Copyright_date = formatDateString(rm_information_item.Copyright_date)
	rm_information_item.Effective_date = formatDateString(rm_information_item.Effective_date)
	rm_information_item.Expiry_date = formatDateString(rm_information_item.Expiry_date)
	rm_information_item.Issue_date = formatDateString(rm_information_item.Issue_date)
	rm_information_item.Modified_date = formatDateString(rm_information_item.Modified_date)
	rm_information_item.Origin_date = formatDateString(rm_information_item.Origin_date)
	rm_information_item.Publish_date = formatDateString(rm_information_item.Publish_date)
	rm_information_item.Submit_date = formatDateString(rm_information_item.Submit_date)
	rm_information_item.Row_effective_date = formatDateString(rm_information_item.Row_effective_date)
	rm_information_item.Row_expiry_date = formatDateString(rm_information_item.Row_expiry_date)






	rows, err := stmt.Exec(rm_information_item.Information_item_id, rm_information_item.Abstract, rm_information_item.Accepted_date, rm_information_item.Access_condition, rm_information_item.Active_ind, rm_information_item.Available_date, rm_information_item.Coord_acquisition_id, rm_information_item.Copyright_date, rm_information_item.Effective_date, rm_information_item.Expiry_date, rm_information_item.Geog_coord_system_id, rm_information_item.Group_ind, rm_information_item.Issue_date, rm_information_item.Item_category, rm_information_item.Item_sub_category, rm_information_item.Local_coord_system_id, rm_information_item.Map_coord_system_id, rm_information_item.Max_latitude, rm_information_item.Max_longitude, rm_information_item.Min_latitude, rm_information_item.Min_longitude, rm_information_item.Modified_date, rm_information_item.Origin_date, rm_information_item.Ppdm_guid, rm_information_item.Publish_date, rm_information_item.Purpose, rm_information_item.Reference_num, rm_information_item.Remark, rm_information_item.Security_desc, rm_information_item.Source, rm_information_item.Source_document_id, rm_information_item.Submit_date, rm_information_item.Time_period_desc, rm_information_item.Title, rm_information_item.Use_condition, rm_information_item.Version_num, rm_information_item.Row_changed_by, rm_information_item.Row_changed_date, rm_information_item.Row_created_by, rm_information_item.Row_effective_date, rm_information_item.Row_expiry_date, rm_information_item.Row_quality, rm_information_item.Info_item_subtype)
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

func PatchRmInformationItem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_information_item set "
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
		} else if key == "accepted_date" || key == "available_date" || key == "copyright_date" || key == "effective_date" || key == "expiry_date" || key == "issue_date" || key == "modified_date" || key == "origin_date" || key == "publish_date" || key == "submit_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
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

func DeleteRmInformationItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_information_item dto.Rm_information_item
	rm_information_item.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_information_item where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_information_item.Info_item_subtype)
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


