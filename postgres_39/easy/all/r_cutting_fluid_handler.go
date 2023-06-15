package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRCuttingFluid(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_cutting_fluid")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_cutting_fluid

	for rows.Next() {
		var r_cutting_fluid dto.R_cutting_fluid
		if err := rows.Scan(&r_cutting_fluid.Cutting_fluid, &r_cutting_fluid.Abbreviation, &r_cutting_fluid.Active_ind, &r_cutting_fluid.Effective_date, &r_cutting_fluid.Expiry_date, &r_cutting_fluid.Long_name, &r_cutting_fluid.Ppdm_guid, &r_cutting_fluid.Remark, &r_cutting_fluid.Short_name, &r_cutting_fluid.Source, &r_cutting_fluid.Row_changed_by, &r_cutting_fluid.Row_changed_date, &r_cutting_fluid.Row_created_by, &r_cutting_fluid.Row_created_date, &r_cutting_fluid.Row_effective_date, &r_cutting_fluid.Row_expiry_date, &r_cutting_fluid.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_cutting_fluid to result
		result = append(result, r_cutting_fluid)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRCuttingFluid(c *fiber.Ctx) error {
	var r_cutting_fluid dto.R_cutting_fluid

	setDefaults(&r_cutting_fluid)

	if err := json.Unmarshal(c.Body(), &r_cutting_fluid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_cutting_fluid values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_cutting_fluid.Row_created_date = formatDate(r_cutting_fluid.Row_created_date)
	r_cutting_fluid.Row_changed_date = formatDate(r_cutting_fluid.Row_changed_date)
	r_cutting_fluid.Effective_date = formatDateString(r_cutting_fluid.Effective_date)
	r_cutting_fluid.Expiry_date = formatDateString(r_cutting_fluid.Expiry_date)
	r_cutting_fluid.Row_effective_date = formatDateString(r_cutting_fluid.Row_effective_date)
	r_cutting_fluid.Row_expiry_date = formatDateString(r_cutting_fluid.Row_expiry_date)






	rows, err := stmt.Exec(r_cutting_fluid.Cutting_fluid, r_cutting_fluid.Abbreviation, r_cutting_fluid.Active_ind, r_cutting_fluid.Effective_date, r_cutting_fluid.Expiry_date, r_cutting_fluid.Long_name, r_cutting_fluid.Ppdm_guid, r_cutting_fluid.Remark, r_cutting_fluid.Short_name, r_cutting_fluid.Source, r_cutting_fluid.Row_changed_by, r_cutting_fluid.Row_changed_date, r_cutting_fluid.Row_created_by, r_cutting_fluid.Row_created_date, r_cutting_fluid.Row_effective_date, r_cutting_fluid.Row_expiry_date, r_cutting_fluid.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRCuttingFluid(c *fiber.Ctx) error {
	var r_cutting_fluid dto.R_cutting_fluid

	setDefaults(&r_cutting_fluid)

	if err := json.Unmarshal(c.Body(), &r_cutting_fluid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_cutting_fluid.Cutting_fluid = id

    if r_cutting_fluid.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_cutting_fluid set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where cutting_fluid = :17")
	if err != nil {
		return err
	}

	r_cutting_fluid.Row_changed_date = formatDate(r_cutting_fluid.Row_changed_date)
	r_cutting_fluid.Effective_date = formatDateString(r_cutting_fluid.Effective_date)
	r_cutting_fluid.Expiry_date = formatDateString(r_cutting_fluid.Expiry_date)
	r_cutting_fluid.Row_effective_date = formatDateString(r_cutting_fluid.Row_effective_date)
	r_cutting_fluid.Row_expiry_date = formatDateString(r_cutting_fluid.Row_expiry_date)






	rows, err := stmt.Exec(r_cutting_fluid.Abbreviation, r_cutting_fluid.Active_ind, r_cutting_fluid.Effective_date, r_cutting_fluid.Expiry_date, r_cutting_fluid.Long_name, r_cutting_fluid.Ppdm_guid, r_cutting_fluid.Remark, r_cutting_fluid.Short_name, r_cutting_fluid.Source, r_cutting_fluid.Row_changed_by, r_cutting_fluid.Row_changed_date, r_cutting_fluid.Row_created_by, r_cutting_fluid.Row_effective_date, r_cutting_fluid.Row_expiry_date, r_cutting_fluid.Row_quality, r_cutting_fluid.Cutting_fluid)
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

func PatchRCuttingFluid(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_cutting_fluid set "
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
	query += " where cutting_fluid = :id"

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

func DeleteRCuttingFluid(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_cutting_fluid dto.R_cutting_fluid
	r_cutting_fluid.Cutting_fluid = id

	stmt, err := db.Prepare("delete from r_cutting_fluid where cutting_fluid = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_cutting_fluid.Cutting_fluid)
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


