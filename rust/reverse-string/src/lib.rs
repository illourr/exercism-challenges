use unicode_segmentation::UnicodeSegmentation;

// reverse takes a string and returns a reversed version of it
pub fn reverse(input: &str) -> String {
    // convert to grapheme clusters iterable, reverse, and collect
    // back into a String collection
    let a: &str = "a";
    input.graphemes(true).rev().collect()
}
