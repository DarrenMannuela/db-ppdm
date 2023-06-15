package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelCongress(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_congress")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_congress

	for rows.Next() {
		var sp_parcel_congress dto.Sp_parcel_congress
		if err := rows.Scan(&sp_parcel_congress.Parcel_congress_id, &sp_parcel_congress.Active_ind, &sp_parcel_congress.Area_id, &sp_parcel_congress.Area_type, &sp_parcel_congress.Congress_range, &sp_parcel_congress.Congress_section, &sp_parcel_congress.Congress_section_suffix, &sp_parcel_congress.Congress_township, &sp_parcel_congress.Congress_twp_name, &sp_parcel_congress.Coord_system_id, &sp_parcel_congress.Description, &sp_parcel_congress.Effective_date, &sp_parcel_congress.Ew_direction, &sp_parcel_congress.Expiry_date, &sp_parcel_congress.Ns_direction, &sp_parcel_congress.Ppdm_guid, &sp_parcel_congress.Principal_meridian, &sp_parcel_congress.Reference_plan_num, &sp_parcel_congress.Remark, &sp_parcel_congress.Section_type, &sp_parcel_congress.Source, &sp_parcel_congress.Spatial_description_id, &sp_parcel_congress.Spatial_obs_no, &sp_parcel_congress.Spot_code, &sp_parcel_congress.Row_changed_by, &sp_parcel_congress.Row_changed_date, &sp_parcel_congress.Row_created_by, &sp_parcel_congress.Row_created_date, &sp_parcel_congress.Row_effective_date, &sp_parcel_congress.Row_expiry_date, &sp_parcel_congress.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_congress to result
		result = append(result, sp_parcel_congress)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelCongress(c *fiber.Ctx) error {
	var sp_parcel_congress dto.Sp_parcel_congress

	setDefaults(&sp_parcel_congress)

	if err := json.Unmarshal(c.Body(), &sp_parcel_congress); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_congress values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	sp_parcel_congress.Row_created_date = formatDate(sp_parcel_congress.Row_created_date)
	sp_parcel_congress.Row_changed_date = formatDate(sp_parcel_congress.Row_changed_date)
	sp_parcel_congress.Effective_date = formatDateString(sp_parcel_congress.Effective_date)
	sp_parcel_congress.Expiry_date = formatDateString(sp_parcel_congress.Expiry_date)
	sp_parcel_congress.Row_effective_date = formatDateString(sp_parcel_congress.Row_effective_date)
	sp_parcel_congress.Row_expiry_date = formatDateString(sp_parcel_congress.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_congress.Parcel_congress_id, sp_parcel_congress.Active_ind, sp_parcel_congress.Area_id, sp_parcel_congress.Area_type, sp_parcel_congress.Congress_range, sp_parcel_congress.Congress_section, sp_parcel_congress.Congress_section_suffix, sp_parcel_congress.Congress_township, sp_parcel_congress.Congress_twp_name, sp_parcel_congress.Coord_system_id, sp_parcel_congress.Description, sp_parcel_congress.Effective_date, sp_parcel_congress.Ew_direction, sp_parcel_congress.Expiry_date, sp_parcel_congress.Ns_direction, sp_parcel_congress.Ppdm_guid, sp_parcel_congress.Principal_meridian, sp_parcel_congress.Reference_plan_num, sp_parcel_congress.Remark, sp_parcel_congress.Section_type, sp_parcel_congress.Source, sp_parcel_congress.Spatial_description_id, sp_parcel_congress.Spatial_obs_no, sp_parcel_congress.Spot_code, sp_parcel_congress.Row_changed_by, sp_parcel_congress.Row_changed_date, sp_parcel_congress.Row_created_by, sp_parcel_congress.Row_created_date, sp_parcel_congress.Row_effective_date, sp_parcel_congress.Row_expiry_date, sp_parcel_congress.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelCongress(c *fiber.Ctx) error {
	var sp_parcel_congress dto.Sp_parcel_congress

	setDefaults(&sp_parcel_congress)

	if err := json.Unmarshal(c.Body(), &sp_parcel_congress); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_congress.Parcel_congress_id = id

    if sp_parcel_congress.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_congress set active_ind = :1, area_id = :2, area_type = :3, congress_range = :4, congress_section = :5, congress_section_suffix = :6, congress_township = :7, congress_twp_name = :8, coord_system_id = :9, description = :10, effective_date = :11, ew_direction = :12, expiry_date = :13, ns_direction = :14, ppdm_guid = :15, principal_meridian = :16, reference_plan_num = :17, remark = :18, section_type = :19, source = :20, spatial_description_id = :21, spatial_obs_no = :22, spot_code = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where parcel_congress_id = :31")
	if err != nil {
		return err
	}

	sp_parcel_congress.Row_changed_date = formatDate(sp_parcel_congress.Row_changed_date)
	sp_parcel_congress.Effective_date = formatDateString(sp_parcel_congress.Effective_date)
	sp_parcel_congress.Expiry_date = formatDateString(sp_parcel_congress.Expiry_date)
	sp_parcel_congress.Row_effective_date = formatDateString(sp_parcel_congress.Row_effective_date)
	sp_parcel_congress.Row_expiry_date = formatDateString(sp_parcel_congress.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_congress.Active_ind, sp_parcel_congress.Area_id, sp_parcel_congress.Area_type, sp_parcel_congress.Congress_range, sp_parcel_congress.Congress_section, sp_parcel_congress.Congress_section_suffix, sp_parcel_congress.Congress_township, sp_parcel_congress.Congress_twp_name, sp_parcel_congress.Coord_system_id, sp_parcel_congress.Description, sp_parcel_congress.Effective_date, sp_parcel_congress.Ew_direction, sp_parcel_congress.Expiry_date, sp_parcel_congress.Ns_direction, sp_parcel_congress.Ppdm_guid, sp_parcel_congress.Principal_meridian, sp_parcel_congress.Reference_plan_num, sp_parcel_congress.Remark, sp_parcel_congress.Section_type, sp_parcel_congress.Source, sp_parcel_congress.Spatial_description_id, sp_parcel_congress.Spatial_obs_no, sp_parcel_congress.Spot_code, sp_parcel_congress.Row_changed_by, sp_parcel_congress.Row_changed_date, sp_parcel_congress.Row_created_by, sp_parcel_congress.Row_effective_date, sp_parcel_congress.Row_expiry_date, sp_parcel_congress.Row_quality, sp_parcel_congress.Parcel_congress_id)
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

func PatchSpParcelCongress(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_congress set "
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
	query += " where parcel_congress_id = :id"

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

func DeleteSpParcelCongress(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_congress dto.Sp_parcel_congress
	sp_parcel_congress.Parcel_congress_id = id

	stmt, err := db.Prepare("delete from sp_parcel_congress where parcel_congress_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_congress.Parcel_congress_id)
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


