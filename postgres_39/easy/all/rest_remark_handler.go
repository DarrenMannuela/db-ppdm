package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rest_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rest_remark

	for rows.Next() {
		var rest_remark dto.Rest_remark
		if err := rows.Scan(&rest_remark.Restriction_id, &rest_remark.Restriction_version, &rest_remark.Restriction_remark_id, &rest_remark.Restriction_remark_seq_no, &rest_remark.Active_ind, &rest_remark.Business_associate_id, &rest_remark.Effective_date, &rest_remark.Expiry_date, &rest_remark.Ppdm_guid, &rest_remark.Remark, &rest_remark.Remark_code, &rest_remark.Remark_date, &rest_remark.Source, &rest_remark.Row_changed_by, &rest_remark.Row_changed_date, &rest_remark.Row_created_by, &rest_remark.Row_created_date, &rest_remark.Row_effective_date, &rest_remark.Row_expiry_date, &rest_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rest_remark to result
		result = append(result, rest_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestRemark(c *fiber.Ctx) error {
	var rest_remark dto.Rest_remark

	setDefaults(&rest_remark)

	if err := json.Unmarshal(c.Body(), &rest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rest_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rest_remark.Row_created_date = formatDate(rest_remark.Row_created_date)
	rest_remark.Row_changed_date = formatDate(rest_remark.Row_changed_date)
	rest_remark.Effective_date = formatDateString(rest_remark.Effective_date)
	rest_remark.Expiry_date = formatDateString(rest_remark.Expiry_date)
	rest_remark.Remark_date = formatDateString(rest_remark.Remark_date)
	rest_remark.Row_effective_date = formatDateString(rest_remark.Row_effective_date)
	rest_remark.Row_expiry_date = formatDateString(rest_remark.Row_expiry_date)






	rows, err := stmt.Exec(rest_remark.Restriction_id, rest_remark.Restriction_version, rest_remark.Restriction_remark_id, rest_remark.Restriction_remark_seq_no, rest_remark.Active_ind, rest_remark.Business_associate_id, rest_remark.Effective_date, rest_remark.Expiry_date, rest_remark.Ppdm_guid, rest_remark.Remark, rest_remark.Remark_code, rest_remark.Remark_date, rest_remark.Source, rest_remark.Row_changed_by, rest_remark.Row_changed_date, rest_remark.Row_created_by, rest_remark.Row_created_date, rest_remark.Row_effective_date, rest_remark.Row_expiry_date, rest_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestRemark(c *fiber.Ctx) error {
	var rest_remark dto.Rest_remark

	setDefaults(&rest_remark)

	if err := json.Unmarshal(c.Body(), &rest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rest_remark.Restriction_id = id

    if rest_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rest_remark set restriction_version = :1, restriction_remark_id = :2, restriction_remark_seq_no = :3, active_ind = :4, business_associate_id = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, remark_code = :10, remark_date = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where restriction_id = :20")
	if err != nil {
		return err
	}

	rest_remark.Row_changed_date = formatDate(rest_remark.Row_changed_date)
	rest_remark.Effective_date = formatDateString(rest_remark.Effective_date)
	rest_remark.Expiry_date = formatDateString(rest_remark.Expiry_date)
	rest_remark.Remark_date = formatDateString(rest_remark.Remark_date)
	rest_remark.Row_effective_date = formatDateString(rest_remark.Row_effective_date)
	rest_remark.Row_expiry_date = formatDateString(rest_remark.Row_expiry_date)






	rows, err := stmt.Exec(rest_remark.Restriction_version, rest_remark.Restriction_remark_id, rest_remark.Restriction_remark_seq_no, rest_remark.Active_ind, rest_remark.Business_associate_id, rest_remark.Effective_date, rest_remark.Expiry_date, rest_remark.Ppdm_guid, rest_remark.Remark, rest_remark.Remark_code, rest_remark.Remark_date, rest_remark.Source, rest_remark.Row_changed_by, rest_remark.Row_changed_date, rest_remark.Row_created_by, rest_remark.Row_effective_date, rest_remark.Row_expiry_date, rest_remark.Row_quality, rest_remark.Restriction_id)
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

func PatchRestRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rest_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where restriction_id = :id"

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

func DeleteRestRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var rest_remark dto.Rest_remark
	rest_remark.Restriction_id = id

	stmt, err := db.Prepare("delete from rest_remark where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rest_remark.Restriction_id)
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


