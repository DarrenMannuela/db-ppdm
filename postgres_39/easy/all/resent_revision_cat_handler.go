package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentRevisionCat(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_revision_cat")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_revision_cat

	for rows.Next() {
		var resent_revision_cat dto.Resent_revision_cat
		if err := rows.Scan(&resent_revision_cat.Revision_category_id, &resent_revision_cat.Active_ind, &resent_revision_cat.Category_name, &resent_revision_cat.Category_type, &resent_revision_cat.Effective_date, &resent_revision_cat.Expiry_date, &resent_revision_cat.Gross_ind, &resent_revision_cat.Net_ind, &resent_revision_cat.Part_of_category_id, &resent_revision_cat.Ppdm_guid, &resent_revision_cat.Remark, &resent_revision_cat.Source, &resent_revision_cat.Row_changed_by, &resent_revision_cat.Row_changed_date, &resent_revision_cat.Row_created_by, &resent_revision_cat.Row_created_date, &resent_revision_cat.Row_effective_date, &resent_revision_cat.Row_expiry_date, &resent_revision_cat.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_revision_cat to result
		result = append(result, resent_revision_cat)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentRevisionCat(c *fiber.Ctx) error {
	var resent_revision_cat dto.Resent_revision_cat

	setDefaults(&resent_revision_cat)

	if err := json.Unmarshal(c.Body(), &resent_revision_cat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_revision_cat values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	resent_revision_cat.Row_created_date = formatDate(resent_revision_cat.Row_created_date)
	resent_revision_cat.Row_changed_date = formatDate(resent_revision_cat.Row_changed_date)
	resent_revision_cat.Effective_date = formatDateString(resent_revision_cat.Effective_date)
	resent_revision_cat.Expiry_date = formatDateString(resent_revision_cat.Expiry_date)
	resent_revision_cat.Row_effective_date = formatDateString(resent_revision_cat.Row_effective_date)
	resent_revision_cat.Row_expiry_date = formatDateString(resent_revision_cat.Row_expiry_date)






	rows, err := stmt.Exec(resent_revision_cat.Revision_category_id, resent_revision_cat.Active_ind, resent_revision_cat.Category_name, resent_revision_cat.Category_type, resent_revision_cat.Effective_date, resent_revision_cat.Expiry_date, resent_revision_cat.Gross_ind, resent_revision_cat.Net_ind, resent_revision_cat.Part_of_category_id, resent_revision_cat.Ppdm_guid, resent_revision_cat.Remark, resent_revision_cat.Source, resent_revision_cat.Row_changed_by, resent_revision_cat.Row_changed_date, resent_revision_cat.Row_created_by, resent_revision_cat.Row_created_date, resent_revision_cat.Row_effective_date, resent_revision_cat.Row_expiry_date, resent_revision_cat.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentRevisionCat(c *fiber.Ctx) error {
	var resent_revision_cat dto.Resent_revision_cat

	setDefaults(&resent_revision_cat)

	if err := json.Unmarshal(c.Body(), &resent_revision_cat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_revision_cat.Revision_category_id = id

    if resent_revision_cat.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_revision_cat set active_ind = :1, category_name = :2, category_type = :3, effective_date = :4, expiry_date = :5, gross_ind = :6, net_ind = :7, part_of_category_id = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where revision_category_id = :19")
	if err != nil {
		return err
	}

	resent_revision_cat.Row_changed_date = formatDate(resent_revision_cat.Row_changed_date)
	resent_revision_cat.Effective_date = formatDateString(resent_revision_cat.Effective_date)
	resent_revision_cat.Expiry_date = formatDateString(resent_revision_cat.Expiry_date)
	resent_revision_cat.Row_effective_date = formatDateString(resent_revision_cat.Row_effective_date)
	resent_revision_cat.Row_expiry_date = formatDateString(resent_revision_cat.Row_expiry_date)






	rows, err := stmt.Exec(resent_revision_cat.Active_ind, resent_revision_cat.Category_name, resent_revision_cat.Category_type, resent_revision_cat.Effective_date, resent_revision_cat.Expiry_date, resent_revision_cat.Gross_ind, resent_revision_cat.Net_ind, resent_revision_cat.Part_of_category_id, resent_revision_cat.Ppdm_guid, resent_revision_cat.Remark, resent_revision_cat.Source, resent_revision_cat.Row_changed_by, resent_revision_cat.Row_changed_date, resent_revision_cat.Row_created_by, resent_revision_cat.Row_effective_date, resent_revision_cat.Row_expiry_date, resent_revision_cat.Row_quality, resent_revision_cat.Revision_category_id)
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

func PatchResentRevisionCat(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_revision_cat set "
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
	query += " where revision_category_id = :id"

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

func DeleteResentRevisionCat(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_revision_cat dto.Resent_revision_cat
	resent_revision_cat.Revision_category_id = id

	stmt, err := db.Prepare("delete from resent_revision_cat where revision_category_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_revision_cat.Revision_category_id)
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


