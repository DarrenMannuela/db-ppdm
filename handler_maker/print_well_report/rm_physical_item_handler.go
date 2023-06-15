package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmPhysicalItem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_physical_item")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_physical_item

	for rows.Next() {
		var rm_physical_item dto.Rm_physical_item
		if err := rows.Scan(&rm_physical_item.Physical_item_id, &rm_physical_item.Active_ind, &rm_physical_item.Bar_code, &rm_physical_item.Catalogue_by_name, &rm_physical_item.Catalogue_date, &rm_physical_item.Certified_true_copy_ind, &rm_physical_item.Circulation_allowed_ind, &rm_physical_item.Circulation_out_ind, &rm_physical_item.Color_format, &rm_physical_item.Create_date, &rm_physical_item.Digital_format, &rm_physical_item.Digital_size, &rm_physical_item.Digital_size_uom, &rm_physical_item.Dim_height, &rm_physical_item.Dim_height_uom, &rm_physical_item.Dim_length, &rm_physical_item.Dim_length_uom, &rm_physical_item.Dim_weight, &rm_physical_item.Dim_weight_uom, &rm_physical_item.Dim_width, &rm_physical_item.Dim_width_uom, &rm_physical_item.Effective_date, &rm_physical_item.Expiry_date, &rm_physical_item.File_count, &rm_physical_item.Group_ind, &rm_physical_item.Image_resolution_uom, &rm_physical_item.Image_x_resolution, &rm_physical_item.Image_y_resolution, &rm_physical_item.Item_category, &rm_physical_item.Item_sub_category, &rm_physical_item.Label, &rm_physical_item.Language, &rm_physical_item.Last_condition, &rm_physical_item.Location_reference, &rm_physical_item.Media_type, &rm_physical_item.Original_file_name, &rm_physical_item.Original_ind, &rm_physical_item.Output_file_name, &rm_physical_item.Page_count, &rm_physical_item.Physical_item_status, &rm_physical_item.Ppdm_guid, &rm_physical_item.Preferred_ind, &rm_physical_item.Qc_by_name, &rm_physical_item.Remark, &rm_physical_item.Retention_period, &rm_physical_item.Sale_allowed_ind, &rm_physical_item.Source, &rm_physical_item.Store_id, &rm_physical_item.Row_changed_by, &rm_physical_item.Row_changed_date, &rm_physical_item.Row_created_by, &rm_physical_item.Row_created_date, &rm_physical_item.Row_effective_date, &rm_physical_item.Row_expiry_date, &rm_physical_item.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_physical_item to result
		result = append(result, rm_physical_item)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmPhysicalItem(c *fiber.Ctx) error {
	var rm_physical_item dto.Rm_physical_item

	setDefaults(&rm_physical_item)

	if err := json.Unmarshal(c.Body(), &rm_physical_item); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_physical_item values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55)")
	if err != nil {
		return err
	}
	rm_physical_item.Row_created_date = formatDate(rm_physical_item.Row_created_date)
	rm_physical_item.Row_changed_date = formatDate(rm_physical_item.Row_changed_date)
	rm_physical_item.Catalogue_date = formatDateString(rm_physical_item.Catalogue_date)
	rm_physical_item.Create_date = formatDateString(rm_physical_item.Create_date)
	rm_physical_item.Effective_date = formatDateString(rm_physical_item.Effective_date)
	rm_physical_item.Expiry_date = formatDateString(rm_physical_item.Expiry_date)
	rm_physical_item.Row_effective_date = formatDateString(rm_physical_item.Row_effective_date)
	rm_physical_item.Row_expiry_date = formatDateString(rm_physical_item.Row_expiry_date)






	rows, err := stmt.Exec(rm_physical_item.Physical_item_id, rm_physical_item.Active_ind, rm_physical_item.Bar_code, rm_physical_item.Catalogue_by_name, rm_physical_item.Catalogue_date, rm_physical_item.Certified_true_copy_ind, rm_physical_item.Circulation_allowed_ind, rm_physical_item.Circulation_out_ind, rm_physical_item.Color_format, rm_physical_item.Create_date, rm_physical_item.Digital_format, rm_physical_item.Digital_size, rm_physical_item.Digital_size_uom, rm_physical_item.Dim_height, rm_physical_item.Dim_height_uom, rm_physical_item.Dim_length, rm_physical_item.Dim_length_uom, rm_physical_item.Dim_weight, rm_physical_item.Dim_weight_uom, rm_physical_item.Dim_width, rm_physical_item.Dim_width_uom, rm_physical_item.Effective_date, rm_physical_item.Expiry_date, rm_physical_item.File_count, rm_physical_item.Group_ind, rm_physical_item.Image_resolution_uom, rm_physical_item.Image_x_resolution, rm_physical_item.Image_y_resolution, rm_physical_item.Item_category, rm_physical_item.Item_sub_category, rm_physical_item.Label, rm_physical_item.Language, rm_physical_item.Last_condition, rm_physical_item.Location_reference, rm_physical_item.Media_type, rm_physical_item.Original_file_name, rm_physical_item.Original_ind, rm_physical_item.Output_file_name, rm_physical_item.Page_count, rm_physical_item.Physical_item_status, rm_physical_item.Ppdm_guid, rm_physical_item.Preferred_ind, rm_physical_item.Qc_by_name, rm_physical_item.Remark, rm_physical_item.Retention_period, rm_physical_item.Sale_allowed_ind, rm_physical_item.Source, rm_physical_item.Store_id, rm_physical_item.Row_changed_by, rm_physical_item.Row_changed_date, rm_physical_item.Row_created_by, rm_physical_item.Row_created_date, rm_physical_item.Row_effective_date, rm_physical_item.Row_expiry_date, rm_physical_item.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmPhysicalItem(c *fiber.Ctx) error {
	var rm_physical_item dto.Rm_physical_item

	setDefaults(&rm_physical_item)

	if err := json.Unmarshal(c.Body(), &rm_physical_item); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_physical_item.Physical_item_id = id

    if rm_physical_item.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_physical_item set active_ind = :1, bar_code = :2, catalogue_by_name = :3, catalogue_date = :4, certified_true_copy_ind = :5, circulation_allowed_ind = :6, circulation_out_ind = :7, color_format = :8, create_date = :9, digital_format = :10, digital_size = :11, digital_size_uom = :12, dim_height = :13, dim_height_uom = :14, dim_length = :15, dim_length_uom = :16, dim_weight = :17, dim_weight_uom = :18, dim_width = :19, dim_width_uom = :20, effective_date = :21, expiry_date = :22, file_count = :23, group_ind = :24, image_resolution_uom = :25, image_x_resolution = :26, image_y_resolution = :27, item_category = :28, item_sub_category = :29, label = :30, language = :31, last_condition = :32, location_reference = :33, media_type = :34, original_file_name = :35, original_ind = :36, output_file_name = :37, page_count = :38, physical_item_status = :39, ppdm_guid = :40, preferred_ind = :41, qc_by_name = :42, remark = :43, retention_period = :44, sale_allowed_ind = :45, source = :46, store_id = :47, row_changed_by = :48, row_changed_date = :49, row_created_by = :50, row_effective_date = :51, row_expiry_date = :52, row_quality = :53 where physical_item_id = :55")
	if err != nil {
		return err
	}

	rm_physical_item.Row_changed_date = formatDate(rm_physical_item.Row_changed_date)
	rm_physical_item.Catalogue_date = formatDateString(rm_physical_item.Catalogue_date)
	rm_physical_item.Create_date = formatDateString(rm_physical_item.Create_date)
	rm_physical_item.Effective_date = formatDateString(rm_physical_item.Effective_date)
	rm_physical_item.Expiry_date = formatDateString(rm_physical_item.Expiry_date)
	rm_physical_item.Row_effective_date = formatDateString(rm_physical_item.Row_effective_date)
	rm_physical_item.Row_expiry_date = formatDateString(rm_physical_item.Row_expiry_date)






	rows, err := stmt.Exec(rm_physical_item.Active_ind, rm_physical_item.Bar_code, rm_physical_item.Catalogue_by_name, rm_physical_item.Catalogue_date, rm_physical_item.Certified_true_copy_ind, rm_physical_item.Circulation_allowed_ind, rm_physical_item.Circulation_out_ind, rm_physical_item.Color_format, rm_physical_item.Create_date, rm_physical_item.Digital_format, rm_physical_item.Digital_size, rm_physical_item.Digital_size_uom, rm_physical_item.Dim_height, rm_physical_item.Dim_height_uom, rm_physical_item.Dim_length, rm_physical_item.Dim_length_uom, rm_physical_item.Dim_weight, rm_physical_item.Dim_weight_uom, rm_physical_item.Dim_width, rm_physical_item.Dim_width_uom, rm_physical_item.Effective_date, rm_physical_item.Expiry_date, rm_physical_item.File_count, rm_physical_item.Group_ind, rm_physical_item.Image_resolution_uom, rm_physical_item.Image_x_resolution, rm_physical_item.Image_y_resolution, rm_physical_item.Item_category, rm_physical_item.Item_sub_category, rm_physical_item.Label, rm_physical_item.Language, rm_physical_item.Last_condition, rm_physical_item.Location_reference, rm_physical_item.Media_type, rm_physical_item.Original_file_name, rm_physical_item.Original_ind, rm_physical_item.Output_file_name, rm_physical_item.Page_count, rm_physical_item.Physical_item_status, rm_physical_item.Ppdm_guid, rm_physical_item.Preferred_ind, rm_physical_item.Qc_by_name, rm_physical_item.Remark, rm_physical_item.Retention_period, rm_physical_item.Sale_allowed_ind, rm_physical_item.Source, rm_physical_item.Store_id, rm_physical_item.Row_changed_by, rm_physical_item.Row_changed_date, rm_physical_item.Row_created_by, rm_physical_item.Row_effective_date, rm_physical_item.Row_expiry_date, rm_physical_item.Row_quality, rm_physical_item.Physical_item_id)
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

func PatchRmPhysicalItem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_physical_item set "
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
		} else if key == "catalogue_date" || key == "create_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where physical_item_id = :id"

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

func DeleteRmPhysicalItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_physical_item dto.Rm_physical_item
	rm_physical_item.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_physical_item where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_physical_item.Physical_item_id)
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


