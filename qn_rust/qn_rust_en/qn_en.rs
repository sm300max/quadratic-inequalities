use std::io;

fn main() {
    let mut antmf = String::new();

    println!("Enter the first coefficient:");

    io::stdin()
        .read_line(&mut antmf)
        .expect("error");

    let a: f64 = antmf.trim().parse().expect("Please enter a valid number");

    if a == 0.0 {
        println!("Quadratic expression cannot have a equal to zero");
        return;
    }

    let mut bntmf = String::new();

    println!("Enter the second coefficient:");

    io::stdin()
        .read_line(&mut bntmf)
        .expect("error");

    let b: f64 = bntmf.trim().parse().expect("Please enter a valid number");

    let mut cntmf = String::new();

    println!("Enter the third coefficient:");

    io::stdin()
        .read_line(&mut cntmf)
        .expect("error");

    let c: f64 = cntmf.trim().parse().expect("Please enter a valid number");

    let mut ismoded = String::new();

    println!("Enter operation sign:");

    io::stdin()
        .read_line(&mut ismoded)
        .expect("error");
    
    let is = ismoded.trim();

    if is != "<" && is != "<=" && is != "=" && is != "=>" && is != ">" {
        println!("Error in sign selection, please try again");
        return;
    }

    let mut question = String::new();

    println!("Is the expression correct?");
    println!("{}x² + {}x + {} {} 0 [y/n]", a, b, c, is);

    io::stdin()
        .read_line(&mut question)
        .expect("error");

    let q = question.trim();

    if q != "y" {
        println!("Please try again");
        println!("You entered {}", q);
        return;
    }

    let upordown: u8;
    if a > 0.0 {
        upordown = 1;
    } else if a < 0.0 {
        upordown = 2;
    } else {
        upordown = 0;
    }
    
    println!("Function: y = {}x² + {}x + {}", a, b, c);
    println!("Quadratic function");
    println!("Graph is a parabola");

    match upordown {
        1 => println!("Branches upward"),
        2 => println!("Branches downward"),
        _ => println!("error")
    }
    println!("Function zeros: {}x² + {}x + {} = 0", a, b, c);

    let oneortwo: u8;
    let x: f64;
    let mut xs: (f64, f64) = (0.0, 0.0);
    let d: f64;
    d = (b*b) - (4.0*a*c);

    if d < 0.0 {
        xs.0 = 0.0;
        xs.1 = 0.0;
        x = 0.0;
        oneortwo = 0;
        println!("D = b² - 4ac");
        println!("D = {}² - 4 * {} * {}", b, a, c);
        println!("D = {}", d);
        println!("D < 0");
        println!("The expression has no roots");
        println!("Solved");
    } else if d == 0.0 {
        oneortwo = 1;
        x = -b / (2.0 * a);
        xs.0 = 0.0;
        xs.1 = 0.0;
        println!("D = b² - 4ac");
        println!("D = {}² - 4 * {} * {}", b, a, c);
        println!("D = {}", d);
        println!("The expression has one root");
        println!("x = -b / 2a");
        println!("x = -{} / 2{}", b, a);
        println!("x = {}", x);
        println!("Solved");
    } else if d > 0.0 {
        oneortwo = 2;
        xs.0 = (-b + d.sqrt())/(2.0*a);
        xs.1 = (-b - d.sqrt())/(2.0*a);
        x = 0.0;
        println!("D = b² - 4ac");
        println!("D = {}² - 4 * {} * {}", b, a, c);
        println!("D = {}", d);
        println!("The expression has two roots");
        println!("x = (-b ± √D) / (2a)");

        if a < 0.0 && b < 0.0 {
            println!("x1 = (-({}) + √{}) / 2 * ({}) )", b, d, a);
            println!("x2 = (-({}) - √{}) / 2 * ({}) )", b, d, a);
        } else if b < 0.0 {
            println!("x1 = (-({}) + √{}) / 2 * {} )", b, d, a);
            println!("x2 = (-({}) - √{}) / 2 * {} )", b, d, a);
        } else if a < 0.0 {
            println!("x1 = ( -{} + √{}) / 2 * ({}) )", b, d, a);
            println!("x2 = ( -{} - √{}) / 2 * ({}) )", b, d, a);
        } else {
            println!("x1 = (-{} + √{}) / 2 * {} )", b, d, a);
            println!("x2 = (-({}) - √{}) / 2 * {} )", b, d, a);
        }
        println!("x1 = {}", xs.0);
        println!("x2 = {}", xs.1);
    } else {
        x = 0.0;
        oneortwo = 0;
    }

    let start: f64;
    let end: f64;

    if xs.0 < xs.1 {
        start = xs.0;
        end = xs.1;
    } else if xs.0 > xs.1 {
        start = xs.1;
        end = xs.0;
    } else {
        start = 0.0;
        end = 0.0;
    }

    if oneortwo == 2 {
        if upordown == 1 {
            if is == ">" || is == "=>" {
                if is == ">" {
                    println!("Answer: (-∞;{})∪({};+∞)", start, end);
                } else if is == "=>" {
                    println!("Answer: (-∞;{}]∪[{};+∞)", start, end);
                }
            }
            if is == "<" || is == "<=" {
                if is == "<" {
                    println!("Answer: ({};{})", start, end);
                }
                if is == "<=" {
                    println!("Answer: [{};{}]", start, end);
                }
            }
        }
        if upordown == 2 {
            if is == ">" || is == "=>" {
                if is == ">" {
                    println!("Answer: ({};{})", start, end)
                }
                if is == "=>" {
                    println!("Answer: [{};{}]", start, end);
                }
            }
            if is == "<" || is == "<=" {
                if is == "<" {
                    println!("Answer: (-∞;{})∪({};+∞)", start, end);
                } else if is == "<=" {
                    println!("Answer: (-∞;{}]∪[{};+∞)", start, end);  
                }               
            }
        }
    } else if oneortwo == 1 {
        if upordown == 1 {
            if is == "=>" || is == "<=" {
                if is == "=>" {
                    println!("Answer: (-∞;+∞)");
                } else if is == "<=" {
                    println!("Answer: No solution (∅)");
                }
            } else if is == ">" || is == "<" {
                if is == ">" {
                    println!("Answer: (-∞;{})∪({};+∞)", x, x);
                } else if is == "<" {
                    println!("Answer: No solution (∅)");
                }
            }
        } else if upordown == 2 {
            if is == "=>" || is == "<=" {
                if is == "=>" {
                    println!("Answer: No solution (∅)");
                } else if is == "<=" {
                    println!("Answer: (-∞;+∞)");
                }
            } else if is == ">" || is == "<" {
                if is == ">" {
                    println!("Answer: No solution (∅)");
                } else if is == "<" {
                    println!("Answer: (-∞;{})∪({};+∞)", x, x);
                }
            }
        }
    }

    let mut exitwaiting = String::new();
    io::stdin()
        .read_line(&mut exitwaiting)
        .expect("error");

    // made by sm300max
}