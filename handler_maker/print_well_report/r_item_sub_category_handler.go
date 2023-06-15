package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRItemSubCategory(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_item_sub_category")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_item_sub_category

	for rows.Next() {
		var r_item_sub_category dto.R_item_sub_category
		if err := rows.Scan(&r_item_sub_category.Item_category, &r_item_sub_category.Item_sub_category, &r_item_sub_category.Abbreviation, &r_item_sub_category.Active_ind, &r_item_sub_category.Effective_date, &r_item_sub_category.Expiry_date, &r_item_sub_category.Long_name, &r_item_sub_category.Ppdm_guid, &r_item_sub_category.Remark, &r_item_sub_category.Retention_period, &r_item_sub_category.Short_name, &r_item_sub_category.Source, &r_item_sub_category.Row_changed_by, &r_item_sub_category.Row_changed_date, &r_item_sub_category.Row_created_by, &r_item_sub_category.Row_created_date, &r_item_sub_category.Row_effective_date, &r_item_sub_category.Row_expiry_date, &r_item_sub_category.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_item_sub_category to result
		result = append(result, r_item_sub_category)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRItemSubCategory(c *fiber.Ctx) error {
	var r_item_sub_category dto.R_item_sub_category

	setDefaults(&r_item_sub_category)

	if err := json.Unmarshal(c.Body(), &r_item_sub_category); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_item_sub_category values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	r_item_sub_category.Row_created_date = formatDate(r_item_sub_category.Row_created_date)
	r_item_sub_category.Row_changed_date = formatDate(r_item_sub_category.Row_changed_date)
	r_item_sub_category.Effective_date = formatDateString(r_item_sub_category.Effective_date)
	r_item_sub_category.Expiry_date = formatDateString(r_item_sub_category.Expiry_date)
	r_item_sub_category.Row_effective_date = formatDateString(r_item_sub_category.Row_effective_date)
	r_item_sub_category.Row_expiry_date = formatDateString(r_item_sub_category.Row_expiry_date)






	rows, err := stmt.Exec(r_item_sub_category.Item_category, r_item_sub_category.Item_sub_category, r_item_sub_category.Abbreviation, r_item_sub_category.Active_ind, r_item_sub_category.Effective_date, r_item_sub_category.Expiry_date, r_item_sub_category.Long_name, r_item_sub_category.Ppdm_guid, r_item_sub_category.Remark, r_item_sub_category.Retention_period, r_item_sub_category.Short_name, r_item_sub_category.Source, r_item_sub_category.Row_changed_by, r_item_sub_category.Row_changed_date, r_item_sub_category.Row_created_by, r_item_sub_category.Row_created_date, r_item_sub_category.Row_effective_date, r_item_sub_category.Row_expiry_date, r_item_sub_category.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRItemSubCategory(c *fiber.Ctx) error {
	var r_item_sub_category dto.R_item_sub_category

	setDefaults(&r_item_sub_category)

	if err := json.Unmarshal(c.Body(), &r_item_sub_category); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_item_sub_category.Item_category = id

    if r_item_sub_category.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_item_sub_category set item_sub_category = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, retention_period = :9, short_name = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where item_category = :19")
	if err != nil {
		return err
	}

	r_item_sub_category.Row_changed_date = formatDate(r_item_sub_category.Row_changed_date)
	r_item_sub_category.Effective_date = formatDateString(r_item_sub_category.Effective_date)
	r_item_sub_category.Expiry_date = formatDateString(r_item_sub_category.Expiry_date)
	r_item_sub_category.Row_effective_date = formatDateString(r_item_sub_category.Row_effective_date)
	r_item_sub_category.Row_expiry_date = formatDateString(r_item_sub_category.Row_expiry_date)






	rows, err := stmt.Exec(r_item_sub_category.Item_sub_category, r_item_sub_category.Abbreviation, r_item_sub_category.Active_ind, r_item_sub_category.Effective_date, r_item_sub_category.Expiry_date, r_item_sub_category.Long_name, r_item_sub_category.Ppdm_guid, r_item_sub_category.Remark, r_item_sub_category.Retention_period, r_item_sub_category.Short_name, r_item_sub_category.Source, r_item_sub_category.Row_changed_by, r_item_sub_category.Row_changed_date, r_item_sub_category.Row_created_by, r_item_sub_category.Row_effective_date, r_item_sub_category.Row_expiry_date, r_item_sub_category.Row_quality, r_item_sub_category.Item_category)
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

func PatchRItemSubCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_item_sub_category set "
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
	query += " where item_category = :id"

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

func DeleteRItemSubCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_item_sub_category dto.R_item_sub_category
	r_item_sub_category.Item_category = id

	stmt, err := db.Prepare("delete from r_item_sub_category where item_category = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_item_sub_category.Item_category)
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


