package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRightRestRem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right_rest_rem")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right_rest_rem

	for rows.Next() {
		var land_right_rest_rem dto.Land_right_rest_rem
		if err := rows.Scan(&land_right_rest_rem.Land_right_subtype, &land_right_rest_rem.Land_right_id, &land_right_rest_rem.Restriction_id, &land_right_rest_rem.Restriction_version, &land_right_rest_rem.Remark_id, &land_right_rest_rem.Remark_seq_no, &land_right_rest_rem.Active_ind, &land_right_rest_rem.Effective_date, &land_right_rest_rem.Expiry_date, &land_right_rest_rem.Ppdm_guid, &land_right_rest_rem.Remark, &land_right_rest_rem.Remark_type, &land_right_rest_rem.Restriction_remark_type, &land_right_rest_rem.Source, &land_right_rest_rem.Row_changed_by, &land_right_rest_rem.Row_changed_date, &land_right_rest_rem.Row_created_by, &land_right_rest_rem.Row_created_date, &land_right_rest_rem.Row_effective_date, &land_right_rest_rem.Row_expiry_date, &land_right_rest_rem.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right_rest_rem to result
		result = append(result, land_right_rest_rem)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRightRestRem(c *fiber.Ctx) error {
	var land_right_rest_rem dto.Land_right_rest_rem

	setDefaults(&land_right_rest_rem)

	if err := json.Unmarshal(c.Body(), &land_right_rest_rem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right_rest_rem values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	land_right_rest_rem.Row_created_date = formatDate(land_right_rest_rem.Row_created_date)
	land_right_rest_rem.Row_changed_date = formatDate(land_right_rest_rem.Row_changed_date)
	land_right_rest_rem.Effective_date = formatDateString(land_right_rest_rem.Effective_date)
	land_right_rest_rem.Expiry_date = formatDateString(land_right_rest_rem.Expiry_date)
	land_right_rest_rem.Row_effective_date = formatDateString(land_right_rest_rem.Row_effective_date)
	land_right_rest_rem.Row_expiry_date = formatDateString(land_right_rest_rem.Row_expiry_date)






	rows, err := stmt.Exec(land_right_rest_rem.Land_right_subtype, land_right_rest_rem.Land_right_id, land_right_rest_rem.Restriction_id, land_right_rest_rem.Restriction_version, land_right_rest_rem.Remark_id, land_right_rest_rem.Remark_seq_no, land_right_rest_rem.Active_ind, land_right_rest_rem.Effective_date, land_right_rest_rem.Expiry_date, land_right_rest_rem.Ppdm_guid, land_right_rest_rem.Remark, land_right_rest_rem.Remark_type, land_right_rest_rem.Restriction_remark_type, land_right_rest_rem.Source, land_right_rest_rem.Row_changed_by, land_right_rest_rem.Row_changed_date, land_right_rest_rem.Row_created_by, land_right_rest_rem.Row_created_date, land_right_rest_rem.Row_effective_date, land_right_rest_rem.Row_expiry_date, land_right_rest_rem.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRightRestRem(c *fiber.Ctx) error {
	var land_right_rest_rem dto.Land_right_rest_rem

	setDefaults(&land_right_rest_rem)

	if err := json.Unmarshal(c.Body(), &land_right_rest_rem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right_rest_rem.Land_right_subtype = id

    if land_right_rest_rem.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right_rest_rem set land_right_id = :1, restriction_id = :2, restriction_version = :3, remark_id = :4, remark_seq_no = :5, active_ind = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, remark_type = :11, restriction_remark_type = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where land_right_subtype = :21")
	if err != nil {
		return err
	}

	land_right_rest_rem.Row_changed_date = formatDate(land_right_rest_rem.Row_changed_date)
	land_right_rest_rem.Effective_date = formatDateString(land_right_rest_rem.Effective_date)
	land_right_rest_rem.Expiry_date = formatDateString(land_right_rest_rem.Expiry_date)
	land_right_rest_rem.Row_effective_date = formatDateString(land_right_rest_rem.Row_effective_date)
	land_right_rest_rem.Row_expiry_date = formatDateString(land_right_rest_rem.Row_expiry_date)






	rows, err := stmt.Exec(land_right_rest_rem.Land_right_id, land_right_rest_rem.Restriction_id, land_right_rest_rem.Restriction_version, land_right_rest_rem.Remark_id, land_right_rest_rem.Remark_seq_no, land_right_rest_rem.Active_ind, land_right_rest_rem.Effective_date, land_right_rest_rem.Expiry_date, land_right_rest_rem.Ppdm_guid, land_right_rest_rem.Remark, land_right_rest_rem.Remark_type, land_right_rest_rem.Restriction_remark_type, land_right_rest_rem.Source, land_right_rest_rem.Row_changed_by, land_right_rest_rem.Row_changed_date, land_right_rest_rem.Row_created_by, land_right_rest_rem.Row_effective_date, land_right_rest_rem.Row_expiry_date, land_right_rest_rem.Row_quality, land_right_rest_rem.Land_right_subtype)
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

func PatchLandRightRestRem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right_rest_rem set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandRightRestRem(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right_rest_rem dto.Land_right_rest_rem
	land_right_rest_rem.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_right_rest_rem where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right_rest_rem.Land_right_subtype)
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


