package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenOperHist(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_oper_hist")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_oper_hist

	for rows.Next() {
		var pden_oper_hist dto.Pden_oper_hist
		if err := rows.Scan(&pden_oper_hist.Pden_subtype, &pden_oper_hist.Pden_id, &pden_oper_hist.Pden_source, &pden_oper_hist.Business_associate_id, &pden_oper_hist.Operator_obs_no, &pden_oper_hist.Active_ind, &pden_oper_hist.Effective_date, &pden_oper_hist.Expiry_date, &pden_oper_hist.Ppdm_guid, &pden_oper_hist.Remark, &pden_oper_hist.Source, &pden_oper_hist.Row_changed_by, &pden_oper_hist.Row_changed_date, &pden_oper_hist.Row_created_by, &pden_oper_hist.Row_created_date, &pden_oper_hist.Row_effective_date, &pden_oper_hist.Row_expiry_date, &pden_oper_hist.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_oper_hist to result
		result = append(result, pden_oper_hist)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenOperHist(c *fiber.Ctx) error {
	var pden_oper_hist dto.Pden_oper_hist

	setDefaults(&pden_oper_hist)

	if err := json.Unmarshal(c.Body(), &pden_oper_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_oper_hist values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	pden_oper_hist.Row_created_date = formatDate(pden_oper_hist.Row_created_date)
	pden_oper_hist.Row_changed_date = formatDate(pden_oper_hist.Row_changed_date)
	pden_oper_hist.Effective_date = formatDateString(pden_oper_hist.Effective_date)
	pden_oper_hist.Expiry_date = formatDateString(pden_oper_hist.Expiry_date)
	pden_oper_hist.Row_effective_date = formatDateString(pden_oper_hist.Row_effective_date)
	pden_oper_hist.Row_expiry_date = formatDateString(pden_oper_hist.Row_expiry_date)






	rows, err := stmt.Exec(pden_oper_hist.Pden_subtype, pden_oper_hist.Pden_id, pden_oper_hist.Pden_source, pden_oper_hist.Business_associate_id, pden_oper_hist.Operator_obs_no, pden_oper_hist.Active_ind, pden_oper_hist.Effective_date, pden_oper_hist.Expiry_date, pden_oper_hist.Ppdm_guid, pden_oper_hist.Remark, pden_oper_hist.Source, pden_oper_hist.Row_changed_by, pden_oper_hist.Row_changed_date, pden_oper_hist.Row_created_by, pden_oper_hist.Row_created_date, pden_oper_hist.Row_effective_date, pden_oper_hist.Row_expiry_date, pden_oper_hist.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenOperHist(c *fiber.Ctx) error {
	var pden_oper_hist dto.Pden_oper_hist

	setDefaults(&pden_oper_hist)

	if err := json.Unmarshal(c.Body(), &pden_oper_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_oper_hist.Pden_subtype = id

    if pden_oper_hist.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_oper_hist set pden_id = :1, pden_source = :2, business_associate_id = :3, operator_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where pden_subtype = :18")
	if err != nil {
		return err
	}

	pden_oper_hist.Row_changed_date = formatDate(pden_oper_hist.Row_changed_date)
	pden_oper_hist.Effective_date = formatDateString(pden_oper_hist.Effective_date)
	pden_oper_hist.Expiry_date = formatDateString(pden_oper_hist.Expiry_date)
	pden_oper_hist.Row_effective_date = formatDateString(pden_oper_hist.Row_effective_date)
	pden_oper_hist.Row_expiry_date = formatDateString(pden_oper_hist.Row_expiry_date)






	rows, err := stmt.Exec(pden_oper_hist.Pden_id, pden_oper_hist.Pden_source, pden_oper_hist.Business_associate_id, pden_oper_hist.Operator_obs_no, pden_oper_hist.Active_ind, pden_oper_hist.Effective_date, pden_oper_hist.Expiry_date, pden_oper_hist.Ppdm_guid, pden_oper_hist.Remark, pden_oper_hist.Source, pden_oper_hist.Row_changed_by, pden_oper_hist.Row_changed_date, pden_oper_hist.Row_created_by, pden_oper_hist.Row_effective_date, pden_oper_hist.Row_expiry_date, pden_oper_hist.Row_quality, pden_oper_hist.Pden_subtype)
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

func PatchPdenOperHist(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_oper_hist set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenOperHist(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_oper_hist dto.Pden_oper_hist
	pden_oper_hist.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_oper_hist where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_oper_hist.Pden_subtype)
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


