package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRWellStatusXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_well_status_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_well_status_xref

	for rows.Next() {
		var r_well_status_xref dto.R_well_status_xref
		if err := rows.Scan(&r_well_status_xref.Status_xref_id, &r_well_status_xref.Status_xref_obs_no, &r_well_status_xref.Active_ind, &r_well_status_xref.Effective_date, &r_well_status_xref.Expiry_date, &r_well_status_xref.Mapping_count, &r_well_status_xref.Ppdm_guid, &r_well_status_xref.Qualifier_value1, &r_well_status_xref.Qualifier_value2, &r_well_status_xref.Remark, &r_well_status_xref.Source, &r_well_status_xref.Status1, &r_well_status_xref.Status2, &r_well_status_xref.Status_qualifier1, &r_well_status_xref.Status_qualifier2, &r_well_status_xref.Status_type1, &r_well_status_xref.Status_type2, &r_well_status_xref.Row_changed_by, &r_well_status_xref.Row_changed_date, &r_well_status_xref.Row_created_by, &r_well_status_xref.Row_created_date, &r_well_status_xref.Row_effective_date, &r_well_status_xref.Row_expiry_date, &r_well_status_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_well_status_xref to result
		result = append(result, r_well_status_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRWellStatusXref(c *fiber.Ctx) error {
	var r_well_status_xref dto.R_well_status_xref

	setDefaults(&r_well_status_xref)

	if err := json.Unmarshal(c.Body(), &r_well_status_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_well_status_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	r_well_status_xref.Row_created_date = formatDate(r_well_status_xref.Row_created_date)
	r_well_status_xref.Row_changed_date = formatDate(r_well_status_xref.Row_changed_date)
	r_well_status_xref.Effective_date = formatDateString(r_well_status_xref.Effective_date)
	r_well_status_xref.Expiry_date = formatDateString(r_well_status_xref.Expiry_date)
	r_well_status_xref.Row_effective_date = formatDateString(r_well_status_xref.Row_effective_date)
	r_well_status_xref.Row_expiry_date = formatDateString(r_well_status_xref.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_xref.Status_xref_id, r_well_status_xref.Status_xref_obs_no, r_well_status_xref.Active_ind, r_well_status_xref.Effective_date, r_well_status_xref.Expiry_date, r_well_status_xref.Mapping_count, r_well_status_xref.Ppdm_guid, r_well_status_xref.Qualifier_value1, r_well_status_xref.Qualifier_value2, r_well_status_xref.Remark, r_well_status_xref.Source, r_well_status_xref.Status1, r_well_status_xref.Status2, r_well_status_xref.Status_qualifier1, r_well_status_xref.Status_qualifier2, r_well_status_xref.Status_type1, r_well_status_xref.Status_type2, r_well_status_xref.Row_changed_by, r_well_status_xref.Row_changed_date, r_well_status_xref.Row_created_by, r_well_status_xref.Row_created_date, r_well_status_xref.Row_effective_date, r_well_status_xref.Row_expiry_date, r_well_status_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRWellStatusXref(c *fiber.Ctx) error {
	var r_well_status_xref dto.R_well_status_xref

	setDefaults(&r_well_status_xref)

	if err := json.Unmarshal(c.Body(), &r_well_status_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_well_status_xref.Status_xref_id = id

    if r_well_status_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_well_status_xref set status_xref_obs_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, mapping_count = :5, ppdm_guid = :6, qualifier_value1 = :7, qualifier_value2 = :8, remark = :9, source = :10, status1 = :11, status2 = :12, status_qualifier1 = :13, status_qualifier2 = :14, status_type1 = :15, status_type2 = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where status_xref_id = :24")
	if err != nil {
		return err
	}

	r_well_status_xref.Row_changed_date = formatDate(r_well_status_xref.Row_changed_date)
	r_well_status_xref.Effective_date = formatDateString(r_well_status_xref.Effective_date)
	r_well_status_xref.Expiry_date = formatDateString(r_well_status_xref.Expiry_date)
	r_well_status_xref.Row_effective_date = formatDateString(r_well_status_xref.Row_effective_date)
	r_well_status_xref.Row_expiry_date = formatDateString(r_well_status_xref.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_xref.Status_xref_obs_no, r_well_status_xref.Active_ind, r_well_status_xref.Effective_date, r_well_status_xref.Expiry_date, r_well_status_xref.Mapping_count, r_well_status_xref.Ppdm_guid, r_well_status_xref.Qualifier_value1, r_well_status_xref.Qualifier_value2, r_well_status_xref.Remark, r_well_status_xref.Source, r_well_status_xref.Status1, r_well_status_xref.Status2, r_well_status_xref.Status_qualifier1, r_well_status_xref.Status_qualifier2, r_well_status_xref.Status_type1, r_well_status_xref.Status_type2, r_well_status_xref.Row_changed_by, r_well_status_xref.Row_changed_date, r_well_status_xref.Row_created_by, r_well_status_xref.Row_effective_date, r_well_status_xref.Row_expiry_date, r_well_status_xref.Row_quality, r_well_status_xref.Status_xref_id)
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

func PatchRWellStatusXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_well_status_xref set "
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
	query += " where status_xref_id = :id"

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

func DeleteRWellStatusXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_well_status_xref dto.R_well_status_xref
	r_well_status_xref.Status_xref_id = id

	stmt, err := db.Prepare("delete from r_well_status_xref where status_xref_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_well_status_xref.Status_xref_id)
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


