package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenProdStringXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_prod_string_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_prod_string_xref

	for rows.Next() {
		var pden_prod_string_xref dto.Pden_prod_string_xref
		if err := rows.Scan(&pden_prod_string_xref.Pden_subtype, &pden_prod_string_xref.Pden_id, &pden_prod_string_xref.Pden_source, &pden_prod_string_xref.Uwi, &pden_prod_string_xref.String_source, &pden_prod_string_xref.String_id, &pden_prod_string_xref.Pden_prs_xref_seq_no, &pden_prod_string_xref.Active_ind, &pden_prod_string_xref.Effective_date, &pden_prod_string_xref.Expiry_date, &pden_prod_string_xref.Ppdm_guid, &pden_prod_string_xref.Remark, &pden_prod_string_xref.Source, &pden_prod_string_xref.Row_changed_by, &pden_prod_string_xref.Row_changed_date, &pden_prod_string_xref.Row_created_by, &pden_prod_string_xref.Row_created_date, &pden_prod_string_xref.Row_effective_date, &pden_prod_string_xref.Row_expiry_date, &pden_prod_string_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_prod_string_xref to result
		result = append(result, pden_prod_string_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenProdStringXref(c *fiber.Ctx) error {
	var pden_prod_string_xref dto.Pden_prod_string_xref

	setDefaults(&pden_prod_string_xref)

	if err := json.Unmarshal(c.Body(), &pden_prod_string_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_prod_string_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	pden_prod_string_xref.Row_created_date = formatDate(pden_prod_string_xref.Row_created_date)
	pden_prod_string_xref.Row_changed_date = formatDate(pden_prod_string_xref.Row_changed_date)
	pden_prod_string_xref.Effective_date = formatDateString(pden_prod_string_xref.Effective_date)
	pden_prod_string_xref.Expiry_date = formatDateString(pden_prod_string_xref.Expiry_date)
	pden_prod_string_xref.Row_effective_date = formatDateString(pden_prod_string_xref.Row_effective_date)
	pden_prod_string_xref.Row_expiry_date = formatDateString(pden_prod_string_xref.Row_expiry_date)






	rows, err := stmt.Exec(pden_prod_string_xref.Pden_subtype, pden_prod_string_xref.Pden_id, pden_prod_string_xref.Pden_source, pden_prod_string_xref.Uwi, pden_prod_string_xref.String_source, pden_prod_string_xref.String_id, pden_prod_string_xref.Pden_prs_xref_seq_no, pden_prod_string_xref.Active_ind, pden_prod_string_xref.Effective_date, pden_prod_string_xref.Expiry_date, pden_prod_string_xref.Ppdm_guid, pden_prod_string_xref.Remark, pden_prod_string_xref.Source, pden_prod_string_xref.Row_changed_by, pden_prod_string_xref.Row_changed_date, pden_prod_string_xref.Row_created_by, pden_prod_string_xref.Row_created_date, pden_prod_string_xref.Row_effective_date, pden_prod_string_xref.Row_expiry_date, pden_prod_string_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenProdStringXref(c *fiber.Ctx) error {
	var pden_prod_string_xref dto.Pden_prod_string_xref

	setDefaults(&pden_prod_string_xref)

	if err := json.Unmarshal(c.Body(), &pden_prod_string_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_prod_string_xref.Pden_subtype = id

    if pden_prod_string_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_prod_string_xref set pden_id = :1, pden_source = :2, uwi = :3, string_source = :4, string_id = :5, pden_prs_xref_seq_no = :6, active_ind = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where pden_subtype = :20")
	if err != nil {
		return err
	}

	pden_prod_string_xref.Row_changed_date = formatDate(pden_prod_string_xref.Row_changed_date)
	pden_prod_string_xref.Effective_date = formatDateString(pden_prod_string_xref.Effective_date)
	pden_prod_string_xref.Expiry_date = formatDateString(pden_prod_string_xref.Expiry_date)
	pden_prod_string_xref.Row_effective_date = formatDateString(pden_prod_string_xref.Row_effective_date)
	pden_prod_string_xref.Row_expiry_date = formatDateString(pden_prod_string_xref.Row_expiry_date)






	rows, err := stmt.Exec(pden_prod_string_xref.Pden_id, pden_prod_string_xref.Pden_source, pden_prod_string_xref.Uwi, pden_prod_string_xref.String_source, pden_prod_string_xref.String_id, pden_prod_string_xref.Pden_prs_xref_seq_no, pden_prod_string_xref.Active_ind, pden_prod_string_xref.Effective_date, pden_prod_string_xref.Expiry_date, pden_prod_string_xref.Ppdm_guid, pden_prod_string_xref.Remark, pden_prod_string_xref.Source, pden_prod_string_xref.Row_changed_by, pden_prod_string_xref.Row_changed_date, pden_prod_string_xref.Row_created_by, pden_prod_string_xref.Row_effective_date, pden_prod_string_xref.Row_expiry_date, pden_prod_string_xref.Row_quality, pden_prod_string_xref.Pden_subtype)
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

func PatchPdenProdStringXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_prod_string_xref set "
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

func DeletePdenProdStringXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_prod_string_xref dto.Pden_prod_string_xref
	pden_prod_string_xref.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_prod_string_xref where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_prod_string_xref.Pden_subtype)
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


