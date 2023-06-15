package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPrLseUnitStrHist(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pr_lse_unit_str_hist")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pr_lse_unit_str_hist

	for rows.Next() {
		var pr_lse_unit_str_hist dto.Pr_lse_unit_str_hist
		if err := rows.Scan(&pr_lse_unit_str_hist.Uwi, &pr_lse_unit_str_hist.Prod_string_source, &pr_lse_unit_str_hist.String_id, &pr_lse_unit_str_hist.Lease_unit_id, &pr_lse_unit_str_hist.Status_id, &pr_lse_unit_str_hist.Active_ind, &pr_lse_unit_str_hist.Effective_date, &pr_lse_unit_str_hist.End_date, &pr_lse_unit_str_hist.Expiry_date, &pr_lse_unit_str_hist.Ppdm_guid, &pr_lse_unit_str_hist.Remark, &pr_lse_unit_str_hist.Source, &pr_lse_unit_str_hist.Start_date, &pr_lse_unit_str_hist.Row_changed_by, &pr_lse_unit_str_hist.Row_changed_date, &pr_lse_unit_str_hist.Row_created_by, &pr_lse_unit_str_hist.Row_created_date, &pr_lse_unit_str_hist.Row_effective_date, &pr_lse_unit_str_hist.Row_expiry_date, &pr_lse_unit_str_hist.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pr_lse_unit_str_hist to result
		result = append(result, pr_lse_unit_str_hist)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPrLseUnitStrHist(c *fiber.Ctx) error {
	var pr_lse_unit_str_hist dto.Pr_lse_unit_str_hist

	setDefaults(&pr_lse_unit_str_hist)

	if err := json.Unmarshal(c.Body(), &pr_lse_unit_str_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pr_lse_unit_str_hist values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	pr_lse_unit_str_hist.Row_created_date = formatDate(pr_lse_unit_str_hist.Row_created_date)
	pr_lse_unit_str_hist.Row_changed_date = formatDate(pr_lse_unit_str_hist.Row_changed_date)
	pr_lse_unit_str_hist.Effective_date = formatDateString(pr_lse_unit_str_hist.Effective_date)
	pr_lse_unit_str_hist.End_date = formatDateString(pr_lse_unit_str_hist.End_date)
	pr_lse_unit_str_hist.Expiry_date = formatDateString(pr_lse_unit_str_hist.Expiry_date)
	pr_lse_unit_str_hist.Start_date = formatDateString(pr_lse_unit_str_hist.Start_date)
	pr_lse_unit_str_hist.Row_effective_date = formatDateString(pr_lse_unit_str_hist.Row_effective_date)
	pr_lse_unit_str_hist.Row_expiry_date = formatDateString(pr_lse_unit_str_hist.Row_expiry_date)






	rows, err := stmt.Exec(pr_lse_unit_str_hist.Uwi, pr_lse_unit_str_hist.Prod_string_source, pr_lse_unit_str_hist.String_id, pr_lse_unit_str_hist.Lease_unit_id, pr_lse_unit_str_hist.Status_id, pr_lse_unit_str_hist.Active_ind, pr_lse_unit_str_hist.Effective_date, pr_lse_unit_str_hist.End_date, pr_lse_unit_str_hist.Expiry_date, pr_lse_unit_str_hist.Ppdm_guid, pr_lse_unit_str_hist.Remark, pr_lse_unit_str_hist.Source, pr_lse_unit_str_hist.Start_date, pr_lse_unit_str_hist.Row_changed_by, pr_lse_unit_str_hist.Row_changed_date, pr_lse_unit_str_hist.Row_created_by, pr_lse_unit_str_hist.Row_created_date, pr_lse_unit_str_hist.Row_effective_date, pr_lse_unit_str_hist.Row_expiry_date, pr_lse_unit_str_hist.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePrLseUnitStrHist(c *fiber.Ctx) error {
	var pr_lse_unit_str_hist dto.Pr_lse_unit_str_hist

	setDefaults(&pr_lse_unit_str_hist)

	if err := json.Unmarshal(c.Body(), &pr_lse_unit_str_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pr_lse_unit_str_hist.Uwi = id

    if pr_lse_unit_str_hist.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pr_lse_unit_str_hist set prod_string_source = :1, string_id = :2, lease_unit_id = :3, status_id = :4, active_ind = :5, effective_date = :6, end_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, start_date = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	pr_lse_unit_str_hist.Row_changed_date = formatDate(pr_lse_unit_str_hist.Row_changed_date)
	pr_lse_unit_str_hist.Effective_date = formatDateString(pr_lse_unit_str_hist.Effective_date)
	pr_lse_unit_str_hist.End_date = formatDateString(pr_lse_unit_str_hist.End_date)
	pr_lse_unit_str_hist.Expiry_date = formatDateString(pr_lse_unit_str_hist.Expiry_date)
	pr_lse_unit_str_hist.Start_date = formatDateString(pr_lse_unit_str_hist.Start_date)
	pr_lse_unit_str_hist.Row_effective_date = formatDateString(pr_lse_unit_str_hist.Row_effective_date)
	pr_lse_unit_str_hist.Row_expiry_date = formatDateString(pr_lse_unit_str_hist.Row_expiry_date)






	rows, err := stmt.Exec(pr_lse_unit_str_hist.Prod_string_source, pr_lse_unit_str_hist.String_id, pr_lse_unit_str_hist.Lease_unit_id, pr_lse_unit_str_hist.Status_id, pr_lse_unit_str_hist.Active_ind, pr_lse_unit_str_hist.Effective_date, pr_lse_unit_str_hist.End_date, pr_lse_unit_str_hist.Expiry_date, pr_lse_unit_str_hist.Ppdm_guid, pr_lse_unit_str_hist.Remark, pr_lse_unit_str_hist.Source, pr_lse_unit_str_hist.Start_date, pr_lse_unit_str_hist.Row_changed_by, pr_lse_unit_str_hist.Row_changed_date, pr_lse_unit_str_hist.Row_created_by, pr_lse_unit_str_hist.Row_effective_date, pr_lse_unit_str_hist.Row_expiry_date, pr_lse_unit_str_hist.Row_quality, pr_lse_unit_str_hist.Uwi)
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

func PatchPrLseUnitStrHist(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pr_lse_unit_str_hist set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeletePrLseUnitStrHist(c *fiber.Ctx) error {
	id := c.Params("id")
	var pr_lse_unit_str_hist dto.Pr_lse_unit_str_hist
	pr_lse_unit_str_hist.Uwi = id

	stmt, err := db.Prepare("delete from pr_lse_unit_str_hist where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pr_lse_unit_str_hist.Uwi)
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


