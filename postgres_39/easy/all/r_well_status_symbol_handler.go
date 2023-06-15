package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRWellStatusSymbol(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_well_status_symbol")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_well_status_symbol

	for rows.Next() {
		var r_well_status_symbol dto.R_well_status_symbol
		if err := rows.Scan(&r_well_status_symbol.Plot_symbol, &r_well_status_symbol.Plot_symbol_obs_no, &r_well_status_symbol.Active_ind, &r_well_status_symbol.Effective_date, &r_well_status_symbol.Expiry_date, &r_well_status_symbol.Ppdm_guid, &r_well_status_symbol.Qualifier_value, &r_well_status_symbol.Remark, &r_well_status_symbol.Source, &r_well_status_symbol.Status, &r_well_status_symbol.Status_facet_count, &r_well_status_symbol.Status_qualifier, &r_well_status_symbol.Status_type, &r_well_status_symbol.Row_changed_by, &r_well_status_symbol.Row_changed_date, &r_well_status_symbol.Row_created_by, &r_well_status_symbol.Row_created_date, &r_well_status_symbol.Row_effective_date, &r_well_status_symbol.Row_expiry_date, &r_well_status_symbol.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_well_status_symbol to result
		result = append(result, r_well_status_symbol)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRWellStatusSymbol(c *fiber.Ctx) error {
	var r_well_status_symbol dto.R_well_status_symbol

	setDefaults(&r_well_status_symbol)

	if err := json.Unmarshal(c.Body(), &r_well_status_symbol); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_well_status_symbol values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	r_well_status_symbol.Row_created_date = formatDate(r_well_status_symbol.Row_created_date)
	r_well_status_symbol.Row_changed_date = formatDate(r_well_status_symbol.Row_changed_date)
	r_well_status_symbol.Effective_date = formatDateString(r_well_status_symbol.Effective_date)
	r_well_status_symbol.Expiry_date = formatDateString(r_well_status_symbol.Expiry_date)
	r_well_status_symbol.Row_effective_date = formatDateString(r_well_status_symbol.Row_effective_date)
	r_well_status_symbol.Row_expiry_date = formatDateString(r_well_status_symbol.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_symbol.Plot_symbol, r_well_status_symbol.Plot_symbol_obs_no, r_well_status_symbol.Active_ind, r_well_status_symbol.Effective_date, r_well_status_symbol.Expiry_date, r_well_status_symbol.Ppdm_guid, r_well_status_symbol.Qualifier_value, r_well_status_symbol.Remark, r_well_status_symbol.Source, r_well_status_symbol.Status, r_well_status_symbol.Status_facet_count, r_well_status_symbol.Status_qualifier, r_well_status_symbol.Status_type, r_well_status_symbol.Row_changed_by, r_well_status_symbol.Row_changed_date, r_well_status_symbol.Row_created_by, r_well_status_symbol.Row_created_date, r_well_status_symbol.Row_effective_date, r_well_status_symbol.Row_expiry_date, r_well_status_symbol.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRWellStatusSymbol(c *fiber.Ctx) error {
	var r_well_status_symbol dto.R_well_status_symbol

	setDefaults(&r_well_status_symbol)

	if err := json.Unmarshal(c.Body(), &r_well_status_symbol); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_well_status_symbol.Plot_symbol = id

    if r_well_status_symbol.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_well_status_symbol set plot_symbol_obs_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, qualifier_value = :6, remark = :7, source = :8, status = :9, status_facet_count = :10, status_qualifier = :11, status_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where plot_symbol = :20")
	if err != nil {
		return err
	}

	r_well_status_symbol.Row_changed_date = formatDate(r_well_status_symbol.Row_changed_date)
	r_well_status_symbol.Effective_date = formatDateString(r_well_status_symbol.Effective_date)
	r_well_status_symbol.Expiry_date = formatDateString(r_well_status_symbol.Expiry_date)
	r_well_status_symbol.Row_effective_date = formatDateString(r_well_status_symbol.Row_effective_date)
	r_well_status_symbol.Row_expiry_date = formatDateString(r_well_status_symbol.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_symbol.Plot_symbol_obs_no, r_well_status_symbol.Active_ind, r_well_status_symbol.Effective_date, r_well_status_symbol.Expiry_date, r_well_status_symbol.Ppdm_guid, r_well_status_symbol.Qualifier_value, r_well_status_symbol.Remark, r_well_status_symbol.Source, r_well_status_symbol.Status, r_well_status_symbol.Status_facet_count, r_well_status_symbol.Status_qualifier, r_well_status_symbol.Status_type, r_well_status_symbol.Row_changed_by, r_well_status_symbol.Row_changed_date, r_well_status_symbol.Row_created_by, r_well_status_symbol.Row_effective_date, r_well_status_symbol.Row_expiry_date, r_well_status_symbol.Row_quality, r_well_status_symbol.Plot_symbol)
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

func PatchRWellStatusSymbol(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_well_status_symbol set "
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
	query += " where plot_symbol = :id"

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

func DeleteRWellStatusSymbol(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_well_status_symbol dto.R_well_status_symbol
	r_well_status_symbol.Plot_symbol = id

	stmt, err := db.Prepare("delete from r_well_status_symbol where plot_symbol = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_well_status_symbol.Plot_symbol)
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


