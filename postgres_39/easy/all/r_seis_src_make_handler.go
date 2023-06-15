package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRSeisSrcMake(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_seis_src_make")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_seis_src_make

	for rows.Next() {
		var r_seis_src_make dto.R_seis_src_make
		if err := rows.Scan(&r_seis_src_make.Seis_src_make, &r_seis_src_make.Abbreviation, &r_seis_src_make.Active_ind, &r_seis_src_make.Effective_date, &r_seis_src_make.Expiry_date, &r_seis_src_make.Long_name, &r_seis_src_make.Ppdm_guid, &r_seis_src_make.Remark, &r_seis_src_make.Short_name, &r_seis_src_make.Source, &r_seis_src_make.Vib_peak_force, &r_seis_src_make.Vib_peak_force_ouom, &r_seis_src_make.Row_changed_by, &r_seis_src_make.Row_changed_date, &r_seis_src_make.Row_created_by, &r_seis_src_make.Row_created_date, &r_seis_src_make.Row_effective_date, &r_seis_src_make.Row_expiry_date, &r_seis_src_make.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_seis_src_make to result
		result = append(result, r_seis_src_make)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRSeisSrcMake(c *fiber.Ctx) error {
	var r_seis_src_make dto.R_seis_src_make

	setDefaults(&r_seis_src_make)

	if err := json.Unmarshal(c.Body(), &r_seis_src_make); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_seis_src_make values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	r_seis_src_make.Row_created_date = formatDate(r_seis_src_make.Row_created_date)
	r_seis_src_make.Row_changed_date = formatDate(r_seis_src_make.Row_changed_date)
	r_seis_src_make.Effective_date = formatDateString(r_seis_src_make.Effective_date)
	r_seis_src_make.Expiry_date = formatDateString(r_seis_src_make.Expiry_date)
	r_seis_src_make.Row_effective_date = formatDateString(r_seis_src_make.Row_effective_date)
	r_seis_src_make.Row_expiry_date = formatDateString(r_seis_src_make.Row_expiry_date)






	rows, err := stmt.Exec(r_seis_src_make.Seis_src_make, r_seis_src_make.Abbreviation, r_seis_src_make.Active_ind, r_seis_src_make.Effective_date, r_seis_src_make.Expiry_date, r_seis_src_make.Long_name, r_seis_src_make.Ppdm_guid, r_seis_src_make.Remark, r_seis_src_make.Short_name, r_seis_src_make.Source, r_seis_src_make.Vib_peak_force, r_seis_src_make.Vib_peak_force_ouom, r_seis_src_make.Row_changed_by, r_seis_src_make.Row_changed_date, r_seis_src_make.Row_created_by, r_seis_src_make.Row_created_date, r_seis_src_make.Row_effective_date, r_seis_src_make.Row_expiry_date, r_seis_src_make.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRSeisSrcMake(c *fiber.Ctx) error {
	var r_seis_src_make dto.R_seis_src_make

	setDefaults(&r_seis_src_make)

	if err := json.Unmarshal(c.Body(), &r_seis_src_make); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_seis_src_make.Seis_src_make = id

    if r_seis_src_make.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_seis_src_make set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, vib_peak_force = :10, vib_peak_force_ouom = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where seis_src_make = :19")
	if err != nil {
		return err
	}

	r_seis_src_make.Row_changed_date = formatDate(r_seis_src_make.Row_changed_date)
	r_seis_src_make.Effective_date = formatDateString(r_seis_src_make.Effective_date)
	r_seis_src_make.Expiry_date = formatDateString(r_seis_src_make.Expiry_date)
	r_seis_src_make.Row_effective_date = formatDateString(r_seis_src_make.Row_effective_date)
	r_seis_src_make.Row_expiry_date = formatDateString(r_seis_src_make.Row_expiry_date)






	rows, err := stmt.Exec(r_seis_src_make.Abbreviation, r_seis_src_make.Active_ind, r_seis_src_make.Effective_date, r_seis_src_make.Expiry_date, r_seis_src_make.Long_name, r_seis_src_make.Ppdm_guid, r_seis_src_make.Remark, r_seis_src_make.Short_name, r_seis_src_make.Source, r_seis_src_make.Vib_peak_force, r_seis_src_make.Vib_peak_force_ouom, r_seis_src_make.Row_changed_by, r_seis_src_make.Row_changed_date, r_seis_src_make.Row_created_by, r_seis_src_make.Row_effective_date, r_seis_src_make.Row_expiry_date, r_seis_src_make.Row_quality, r_seis_src_make.Seis_src_make)
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

func PatchRSeisSrcMake(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_seis_src_make set "
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
	query += " where seis_src_make = :id"

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

func DeleteRSeisSrcMake(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_seis_src_make dto.R_seis_src_make
	r_seis_src_make.Seis_src_make = id

	stmt, err := db.Prepare("delete from r_seis_src_make where seis_src_make = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_seis_src_make.Seis_src_make)
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


