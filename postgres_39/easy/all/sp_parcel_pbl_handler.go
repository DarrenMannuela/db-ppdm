package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelPbl(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_pbl")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_pbl

	for rows.Next() {
		var sp_parcel_pbl dto.Sp_parcel_pbl
		if err := rows.Scan(&sp_parcel_pbl.Parcel_pbl_id, &sp_parcel_pbl.Active_ind, &sp_parcel_pbl.Area_id, &sp_parcel_pbl.Area_type, &sp_parcel_pbl.Coord_system_id, &sp_parcel_pbl.Description, &sp_parcel_pbl.Effective_date, &sp_parcel_pbl.Expiry_date, &sp_parcel_pbl.Jurisdiction, &sp_parcel_pbl.Pbl_block, &sp_parcel_pbl.Pbl_lot, &sp_parcel_pbl.Pbl_plan, &sp_parcel_pbl.Pbl_subdivision, &sp_parcel_pbl.Ppdm_guid, &sp_parcel_pbl.Remark, &sp_parcel_pbl.Source, &sp_parcel_pbl.Spatial_description_id, &sp_parcel_pbl.Spatial_obs_no, &sp_parcel_pbl.Row_changed_by, &sp_parcel_pbl.Row_changed_date, &sp_parcel_pbl.Row_created_by, &sp_parcel_pbl.Row_created_date, &sp_parcel_pbl.Row_effective_date, &sp_parcel_pbl.Row_expiry_date, &sp_parcel_pbl.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_pbl to result
		result = append(result, sp_parcel_pbl)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelPbl(c *fiber.Ctx) error {
	var sp_parcel_pbl dto.Sp_parcel_pbl

	setDefaults(&sp_parcel_pbl)

	if err := json.Unmarshal(c.Body(), &sp_parcel_pbl); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_pbl values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	sp_parcel_pbl.Row_created_date = formatDate(sp_parcel_pbl.Row_created_date)
	sp_parcel_pbl.Row_changed_date = formatDate(sp_parcel_pbl.Row_changed_date)
	sp_parcel_pbl.Effective_date = formatDateString(sp_parcel_pbl.Effective_date)
	sp_parcel_pbl.Expiry_date = formatDateString(sp_parcel_pbl.Expiry_date)
	sp_parcel_pbl.Row_effective_date = formatDateString(sp_parcel_pbl.Row_effective_date)
	sp_parcel_pbl.Row_expiry_date = formatDateString(sp_parcel_pbl.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_pbl.Parcel_pbl_id, sp_parcel_pbl.Active_ind, sp_parcel_pbl.Area_id, sp_parcel_pbl.Area_type, sp_parcel_pbl.Coord_system_id, sp_parcel_pbl.Description, sp_parcel_pbl.Effective_date, sp_parcel_pbl.Expiry_date, sp_parcel_pbl.Jurisdiction, sp_parcel_pbl.Pbl_block, sp_parcel_pbl.Pbl_lot, sp_parcel_pbl.Pbl_plan, sp_parcel_pbl.Pbl_subdivision, sp_parcel_pbl.Ppdm_guid, sp_parcel_pbl.Remark, sp_parcel_pbl.Source, sp_parcel_pbl.Spatial_description_id, sp_parcel_pbl.Spatial_obs_no, sp_parcel_pbl.Row_changed_by, sp_parcel_pbl.Row_changed_date, sp_parcel_pbl.Row_created_by, sp_parcel_pbl.Row_created_date, sp_parcel_pbl.Row_effective_date, sp_parcel_pbl.Row_expiry_date, sp_parcel_pbl.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelPbl(c *fiber.Ctx) error {
	var sp_parcel_pbl dto.Sp_parcel_pbl

	setDefaults(&sp_parcel_pbl)

	if err := json.Unmarshal(c.Body(), &sp_parcel_pbl); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_pbl.Parcel_pbl_id = id

    if sp_parcel_pbl.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_pbl set active_ind = :1, area_id = :2, area_type = :3, coord_system_id = :4, description = :5, effective_date = :6, expiry_date = :7, jurisdiction = :8, pbl_block = :9, pbl_lot = :10, pbl_plan = :11, pbl_subdivision = :12, ppdm_guid = :13, remark = :14, source = :15, spatial_description_id = :16, spatial_obs_no = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where parcel_pbl_id = :25")
	if err != nil {
		return err
	}

	sp_parcel_pbl.Row_changed_date = formatDate(sp_parcel_pbl.Row_changed_date)
	sp_parcel_pbl.Effective_date = formatDateString(sp_parcel_pbl.Effective_date)
	sp_parcel_pbl.Expiry_date = formatDateString(sp_parcel_pbl.Expiry_date)
	sp_parcel_pbl.Row_effective_date = formatDateString(sp_parcel_pbl.Row_effective_date)
	sp_parcel_pbl.Row_expiry_date = formatDateString(sp_parcel_pbl.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_pbl.Active_ind, sp_parcel_pbl.Area_id, sp_parcel_pbl.Area_type, sp_parcel_pbl.Coord_system_id, sp_parcel_pbl.Description, sp_parcel_pbl.Effective_date, sp_parcel_pbl.Expiry_date, sp_parcel_pbl.Jurisdiction, sp_parcel_pbl.Pbl_block, sp_parcel_pbl.Pbl_lot, sp_parcel_pbl.Pbl_plan, sp_parcel_pbl.Pbl_subdivision, sp_parcel_pbl.Ppdm_guid, sp_parcel_pbl.Remark, sp_parcel_pbl.Source, sp_parcel_pbl.Spatial_description_id, sp_parcel_pbl.Spatial_obs_no, sp_parcel_pbl.Row_changed_by, sp_parcel_pbl.Row_changed_date, sp_parcel_pbl.Row_created_by, sp_parcel_pbl.Row_effective_date, sp_parcel_pbl.Row_expiry_date, sp_parcel_pbl.Row_quality, sp_parcel_pbl.Parcel_pbl_id)
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

func PatchSpParcelPbl(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_pbl set "
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
	query += " where parcel_pbl_id = :id"

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

func DeleteSpParcelPbl(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_pbl dto.Sp_parcel_pbl
	sp_parcel_pbl.Parcel_pbl_id = id

	stmt, err := db.Prepare("delete from sp_parcel_pbl where parcel_pbl_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_pbl.Parcel_pbl_id)
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


