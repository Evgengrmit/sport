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

$args = array(
    'taxonomy' => 'days',
);
$terms = get_terms($args);
$count = 1;

$results = array();
foreach ($terms as $key => $term) {
    if ($count > 7) break;
    $count++;

    $args = array(
        'numberposts' => 3,
        'category' => $term->term_id,
        'post_type' => 'trainings',
        'tax_query' => array(
            array(
                'taxonomy' => 'days',
                'terms' => $term->term_id,
                'field' => 'id',
            )
        ),
    );

    $query = new WP_query($args);

    while ($query->have_posts()) {
        $query->the_post();
        $wodOrigDescription = get_the_content();
        $wodOrigDescription = html_entity_decode($wodOrigDescription);
        $wodOrigDescription = strip_tags($wodOrigDescription, '<p><strong><br>');
        $wodOrigDescription = str_replace(["\r", "\n"], "", $wodOrigDescription);
        $wodDescription = str_replace('<br /></p>', '</p>', $wodOrigDescription);
        $wodDescription = preg_replace('#(<br\s?/>)+?</p>#isxu', '</p>', $wodDescription);
        $wodDescription = preg_replace('#<p><br\s?/>#isxu', '<p>', $wodDescription);
        $wodDescription = preg_replace("#<p>\s+</p>#isxu", '', $wodDescription);
        $wodDescription = preg_replace("#^<p>(.*?)</p>$#isxu", '<span>$1</span>', $wodDescription);

        $title = get_the_title();

        $results[] = [
            'title' => $title,
            'scheduledAt' => $term->name,
            'description' => $wodDescription,
            'origDescription' => $wodOrigDescription,
        ];
    }
}

header("Content-Type: application/json");
echo json_encode($results, JSON_UNESCAPED_UNICODE);