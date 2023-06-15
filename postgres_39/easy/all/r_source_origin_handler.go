package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRSourceOrigin(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_source_origin")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_source_origin

	for rows.Next() {
		var r_source_origin dto.R_source_origin
		if err := rows.Scan(&r_source_origin.Source, &r_source_origin.Origin_obs_no, &r_source_origin.Abbreviation, &r_source_origin.Active_ind, &r_source_origin.Effective_date, &r_source_origin.Expiry_date, &r_source_origin.Long_name, &r_source_origin.Owner_ba_id, &r_source_origin.Physical_item_id, &r_source_origin.Ppdm_guid, &r_source_origin.Project_id, &r_source_origin.Remark, &r_source_origin.Row_source, &r_source_origin.Short_name, &r_source_origin.Source_document, &r_source_origin.Sw_application_id, &r_source_origin.Row_changed_by, &r_source_origin.Row_changed_date, &r_source_origin.Row_created_by, &r_source_origin.Row_created_date, &r_source_origin.Row_effective_date, &r_source_origin.Row_expiry_date, &r_source_origin.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_source_origin to result
		result = append(result, r_source_origin)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRSourceOrigin(c *fiber.Ctx) error {
	var r_source_origin dto.R_source_origin

	setDefaults(&r_source_origin)

	if err := json.Unmarshal(c.Body(), &r_source_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_source_origin values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	r_source_origin.Row_created_date = formatDate(r_source_origin.Row_created_date)
	r_source_origin.Row_changed_date = formatDate(r_source_origin.Row_changed_date)
	r_source_origin.Effective_date = formatDateString(r_source_origin.Effective_date)
	r_source_origin.Expiry_date = formatDateString(r_source_origin.Expiry_date)
	r_source_origin.Row_effective_date = formatDateString(r_source_origin.Row_effective_date)
	r_source_origin.Row_expiry_date = formatDateString(r_source_origin.Row_expiry_date)






	rows, err := stmt.Exec(r_source_origin.Source, r_source_origin.Origin_obs_no, r_source_origin.Abbreviation, r_source_origin.Active_ind, r_source_origin.Effective_date, r_source_origin.Expiry_date, r_source_origin.Long_name, r_source_origin.Owner_ba_id, r_source_origin.Physical_item_id, r_source_origin.Ppdm_guid, r_source_origin.Project_id, r_source_origin.Remark, r_source_origin.Row_source, r_source_origin.Short_name, r_source_origin.Source_document, r_source_origin.Sw_application_id, r_source_origin.Row_changed_by, r_source_origin.Row_changed_date, r_source_origin.Row_created_by, r_source_origin.Row_created_date, r_source_origin.Row_effective_date, r_source_origin.Row_expiry_date, r_source_origin.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRSourceOrigin(c *fiber.Ctx) error {
	var r_source_origin dto.R_source_origin

	setDefaults(&r_source_origin)

	if err := json.Unmarshal(c.Body(), &r_source_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_source_origin.Source = id

    if r_source_origin.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_source_origin set origin_obs_no = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, owner_ba_id = :7, physical_item_id = :8, ppdm_guid = :9, project_id = :10, remark = :11, row_source = :12, short_name = :13, source_document = :14, sw_application_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where source = :23")
	if err != nil {
		return err
	}

	r_source_origin.Row_changed_date = formatDate(r_source_origin.Row_changed_date)
	r_source_origin.Effective_date = formatDateString(r_source_origin.Effective_date)
	r_source_origin.Expiry_date = formatDateString(r_source_origin.Expiry_date)
	r_source_origin.Row_effective_date = formatDateString(r_source_origin.Row_effective_date)
	r_source_origin.Row_expiry_date = formatDateString(r_source_origin.Row_expiry_date)






	rows, err := stmt.Exec(r_source_origin.Origin_obs_no, r_source_origin.Abbreviation, r_source_origin.Active_ind, r_source_origin.Effective_date, r_source_origin.Expiry_date, r_source_origin.Long_name, r_source_origin.Owner_ba_id, r_source_origin.Physical_item_id, r_source_origin.Ppdm_guid, r_source_origin.Project_id, r_source_origin.Remark, r_source_origin.Row_source, r_source_origin.Short_name, r_source_origin.Source_document, r_source_origin.Sw_application_id, r_source_origin.Row_changed_by, r_source_origin.Row_changed_date, r_source_origin.Row_created_by, r_source_origin.Row_effective_date, r_source_origin.Row_expiry_date, r_source_origin.Row_quality, r_source_origin.Source)
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

func PatchRSourceOrigin(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_source_origin set "
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
	query += " where source = :id"

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

func DeleteRSourceOrigin(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_source_origin dto.R_source_origin
	r_source_origin.Source = id

	stmt, err := db.Prepare("delete from r_source_origin where source = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_source_origin.Source)
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


