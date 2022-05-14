<?php
/**
 * Front to the WordPress application. This file doesn't do anything, but loads
 * wp-blog-header.php which does and tells WordPress to load the theme.
 *
 * @package WordPress
 */

/**
 * Tells WordPress to load the WordPress theme and output it.
 *
 * @var bool
 */
define('WP_USE_THEMES', false);

/** Loads the WordPress Environment and Template */
require(dirname(__FILE__) . '/wp-blog-header.php');

// Load the WordPress library.
require_once(dirname(__FILE__) . '/wp-load.php');

// Set up the WordPress query.
wp();

$pages = get_pages(array(
    'meta_key' => '_wp_page_template',
    'meta_value' => 'page-table.php'
));
$id = $pages[0]->ID;
$table = get_field('строка', $id);
//print_r($table);


$schedule = get_posts(array(
    'numberposts' => -1,
    'orderby' => 'date',
    'order' => 'ASC',
    'post_type' => 'timetable'
));

$result = [];
foreach ($schedule as $s_time) {
    $s_days = get_fields($s_time->ID);
    for ($x = 0; $x++ < 7;) {
        $show_key = 'show_' . $x;
        $tranings_key = 'tranings_' . $x;


        if ($s_days[$show_key] == '1') {
            foreach ($s_days[$tranings_key] as $train) {
                $trainerPic = get_field('фото', $train['тренер']->ID);
                if (empty($trainerPic)) {
                    $trainerPic = null;
                }
                $result[] = [
                    'day' => $x,
                    'time' => $s_time->post_title,
                    'name' => $train['name'],
                    'trainerName' => get_field('имя', $train['тренер']->ID),
                    'trainerPic' => $trainerPic,
                ];
            }
        }
    }
}

header("Content-Type: application/json");
echo json_encode($result);